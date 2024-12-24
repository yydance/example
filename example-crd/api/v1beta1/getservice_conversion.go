package v1beta1

import (
	v1alpha1 "example.cn/api/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

func (src *GetService) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha1.GetService)
	dst.ObjectMeta = src.ObjectMeta
	dst.Spec.GetAll = src.Spec.GetAll
	dst.Spec.MatchStr = src.Spec.MatchStr
	dst.Spec.Regex = src.Spec.Regex
	dst.Spec.Namespace = src.Spec.Namespace
	dst.Status.Complated = src.Status.Completed
	dst.Status.Status = src.Status.Status

	return nil
}

func (dst *GetService) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha1.GetService)
	dst.ObjectMeta = src.ObjectMeta
	dst.Spec.GetAll = src.Spec.GetAll
	dst.Spec.MatchStr = src.Spec.MatchStr
	dst.Spec.Regex = src.Spec.Regex
	dst.Spec.Namespace = src.Spec.Namespace
	dst.Status.Completed = src.Status.Complated

	return nil
}
