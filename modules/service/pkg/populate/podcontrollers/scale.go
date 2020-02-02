package podcontrollers

//
////
//type scaleParams struct {
//	Scale          *int32
//	MaxSurge       *intstr.IntOrString
//	MaxUnavailable *intstr.IntOrString
//}
//
//func parseScaleParams(w riov1.Wrangler) scaleParams {
//	var scale *int
//	scale = w.GetSpec().Replicas
//
//	if w.GetStatus().ComputedReplicas != nil && services.AutoscaleEnable(w) {
//		scale = w.GetStatus().ComputedReplicas
//	}
//
//	// at one point we told users that -1 meant we don't control scale. nil is now that behavior
//	if scale != nil && *scale < 0 {
//		scale = nil
//	}
//
//	sp := scaleParams{
//		MaxSurge:       w.GetSpec().MaxSurge,
//		MaxUnavailable: w.GetSpec().MaxUnavailable,
//	}
//
//	if scale != nil {
//		scale32 := int32(*scale)
//		sp.Scale = &scale32
//	}
//
//	return sp
//}
