package arrgo

func Where(cond *Arrb, tv, fv interface{}) *Arrf {
	t := Zeros(cond.shape...)
    for i, v := range cond.data {
        if v {
            switch tv.(type) {
            case float64:
                t.data[i] = tv.(float64)
            case float32:
                t.data[i] = float64(tv.(float32))
            case int:
                t.data[i] = float64(tv.(int))
            case *Arrf:
                t.data[i] = tv.(*Arrf).data[i]
            default:
                panic(TYPE_ERROR)
            }
        } else {
            switch fv.(type) {
            case float64:
                t.data[i] = fv.(float64)
            case float32:
                t.data[i] = float64(fv.(float32))
            case int:
                t.data[i] = float64(fv.(int))
            case *Arrf:
                t.data[i] = fv.(*Arrf).data[i]
            default:
                panic(TYPE_ERROR)
            }
        }
    }
    return t
}
