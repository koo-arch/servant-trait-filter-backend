package validation

import (
    "fmt"

    "github.com/koo-arch/servant-trait-filter-backend/internal/httperr"
	"github.com/koo-arch/servant-trait-filter-backend/internal/search"
)

const (
    maxDepth       = 8   // ネスト許容深さ（調整可）
    maxTotalNodes  = 500 // ノード総数の上限（DoS 対策）
)

type exprStats struct {
    depthMax int
    total    int
}

func ValidateSearchRequest(req *search.ServantSearchQuery) error {
    // limit/offset は gin の binding が弾く

    var fields []httperr.FieldError
    st := &exprStats{}
    validateExpr(&req.Root, "root", 1, st, &fields)

    // 深さ・ノード数
    if st.depthMax > maxDepth {
        fields = append(fields, fe("root", fmt.Sprintf("expression depth exceeds max(%d)", maxDepth)))
    }
    if st.total > maxTotalNodes {
        fields = append(fields, fe("root", fmt.Sprintf("expression size exceeds max(%d)", maxTotalNodes)))
    }

    if len(fields) > 0 {
        // 共通エラーフォーマットで返す
        return &httperr.AppError{
            Code:       httperr.CodeValidationFailed,
            HTTPStatus: 400,
            Message:    "入力値が不正です。",
            Details:    fields, // httperr.Write 側で fieldErrors に入れる前提
        }
    }
    return nil
}

func validateExpr(e *search.Expr, path string, depth int, st *exprStats, out *[]httperr.FieldError) {
    if e == nil {
        *out = append(*out, fe(path, "must not be null"))
        return
    }
    if depth > st.depthMax {
        st.depthMax = depth
    }
    st.total++

    // 子の存在チェック
    hasAnd := len(e.And) > 0
    hasOr := len(e.Or) > 0
    hasNot := e.Not != nil

    // 原子（ID が1つでも指定されているか）
    atomCount := countAtoms(e)

    // AND/OR はどちらか一方のみ許可
    if hasAnd && hasOr {
        *out = append(*out, fe(path, "must not have both 'and' and 'or'"))
    }

    // ルートが完全に空（子も原子もなし）は 全ての概念を許容するのでOK
    if !hasAnd && !hasOr && !hasNot && atomCount == 0 {
        return
    }

    // Not は1つだけ（現在の設計では Not.* の下にさらに And/Or/Not/atoms を置けるが、NULL は不可）
    if hasNot {
        validateExpr(e.Not, path+".not", depth+1, st, out)
    }

    // And/Or の子を再帰検証
    for i, c := range e.And {
        validateExpr(c, fmt.Sprintf("%s.and[%d]", path, i), depth+1, st, out)
    }
    for i, c := range e.Or {
        validateExpr(c, fmt.Sprintf("%s.or[%d]", path, i), depth+1, st, out)
    }

    // 原子（ID は正整数）
    checkPositive(e.TraitID, path+".trait", out)
    checkPositive(e.ClassID, path+".class", out)
    checkPositive(e.AttributeID, path+".attribute", out)
    checkPositive(e.OrderAlignID, path+".orderAlignment", out)
    checkPositive(e.MoralAlignID, path+".moralAlignment", out)

    // 同一ノードに複数概念の混在禁止
    if atomCount > 0 && (hasAnd || hasOr || hasNot) {
        *out = append(*out, fe(path, "must not mix atoms with logical operators"))
    }
}

func countAtoms(e *search.Expr) int {
    n := 0
    if e.TraitID != nil { n++ }
    if e.ClassID != nil { n++ }
    if e.AttributeID != nil { n++ }
    if e.OrderAlignID != nil { n++ }
    if e.MoralAlignID != nil { n++ }
    return n
}

func checkPositive(p *int, field string, out *[]httperr.FieldError) {
    if p == nil { return }
    if *p <= 0 {
        *out = append(*out, fe(field, "must be positive integer"))
    }
}

func fe(field, reason string) httperr.FieldError {
    return httperr.FieldError{
        Field:  field,
        Reason: reason,
    }
}