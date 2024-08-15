package k8stest

import (
	errors "github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
)

// PodTemplateSpec holds the api's podtemplatespec objects
type PodTemplateSpec struct {
	Object *corev1.PodTemplateSpec
}

// Builder is the builder object for Pod
type PodtemplatespecBuilder struct {
	podtemplatespec *PodTemplateSpec
	errs            []error
}

// NewBuilder returns new instance of Builder
func NewPodtemplatespecBuilder() *PodtemplatespecBuilder {
	return &PodtemplatespecBuilder{
		podtemplatespec: &PodTemplateSpec{
			Object: &corev1.PodTemplateSpec{},
		},
	}
}

// WithName sets the Name field of podtemplatespec with provided value.
func (b *PodtemplatespecBuilder) WithName(name string) *PodtemplatespecBuilder {
	if len(name) == 0 {
		b.errs = append(
			b.errs,
			errors.New("failed to build podtemplatespec object: missing name"),
		)
		return b
	}
	b.podtemplatespec.Object.Name = name
	return b
}

// WithNamespace sets the Namespace field of PodTemplateSpec with provided value.
func (b *PodtemplatespecBuilder) WithNamespace(namespace string) *PodtemplatespecBuilder {
	if len(namespace) == 0 {
		b.errs = append(
			b.errs,
			errors.New(
				"failed to build podtemplatespec object: missing namespace",
			),
		)
		return b
	}
	b.podtemplatespec.Object.Namespace = namespace
	return b
}

// WithAnnotations merges existing annotations if any
// with the ones that are provided here
func (b *PodtemplatespecBuilder) WithAnnotations(annotations map[string]string) *PodtemplatespecBuilder {
	if len(annotations) == 0 {
		b.errs = append(
			b.errs,
			errors.New("failed to build deployment object: missing annotations"),
		)
		return b
	}

	if b.podtemplatespec.Object.Annotations == nil {
		return b.WithAnnotationsNew(annotations)
	}

	for key, value := range annotations {
		b.podtemplatespec.Object.Annotations[key] = value
	}
	return b
}

// WithAnnotationsNew resets the annotation field of podtemplatespec
// with provided arguments
func (b *PodtemplatespecBuilder) WithAnnotationsNew(annotations map[string]string) *PodtemplatespecBuilder {
	if len(annotations) == 0 {
		b.errs = append(
			b.errs,
			errors.New(
				"failed to build podtemplatespec object: missing annotations",
			),
		)
		return b
	}

	// copy of original map
	newannotations := map[string]string{}
	for key, value := range annotations {
		newannotations[key] = value
	}

	// override
	b.podtemplatespec.Object.Annotations = newannotations
	return b
}

// WithLabels merges existing labels if any
// with the ones that are provided here
func (b *PodtemplatespecBuilder) WithLabels(labels map[string]string) *PodtemplatespecBuilder {
	if len(labels) == 0 {
		b.errs = append(
			b.errs,
			errors.New(
				"failed to build podtemplatespec object: missing labels",
			),
		)
		return b
	}

	if b.podtemplatespec.Object.Labels == nil {
		return b.WithLabelsNew(labels)
	}

	for key, value := range labels {
		b.podtemplatespec.Object.Labels[key] = value
	}
	return b
}

// WithLabelsNew resets the labels field of podtemplatespec
// with provided arguments
func (b *PodtemplatespecBuilder) WithLabelsNew(labels map[string]string) *PodtemplatespecBuilder {
	if len(labels) == 0 {
		b.errs = append(
			b.errs,
			errors.New(
				"failed to build podtemplatespec object: missing labels",
			),
		)
		return b
	}

	// copy of original map
	newlbls := map[string]string{}
	for key, value := range labels {
		newlbls[key] = value
	}

	// override
	b.podtemplatespec.Object.Labels = newlbls
	return b
}

// WithNodeSelector merges the nodeselectors if present
// with the provided arguments
func (b *PodtemplatespecBuilder) WithNodeSelector(nodeselectors map[string]string) *PodtemplatespecBuilder {
	if len(nodeselectors) == 0 {
		b.errs = append(
			b.errs,
			errors.New(
				"failed to build podtemplatespec object: missing nodeselectors",
			),
		)
		return b
	}

	if b.podtemplatespec.Object.Spec.NodeSelector == nil {
		return b.WithNodeSelectorNew(nodeselectors)
	}

	for key, value := range nodeselectors {
		b.podtemplatespec.Object.Spec.NodeSelector[key] = value
	}
	return b
}

// WithPriorityClassName sets the PriorityClassName field of podtemplatespec
func (b *PodtemplatespecBuilder) WithPriorityClassName(prorityClassName string) *PodtemplatespecBuilder {
	b.podtemplatespec.Object.Spec.PriorityClassName = prorityClassName
	return b
}

// WithNodeSelectorNew resets the nodeselector field of podtemplatespec
// with provided arguments
func (b *PodtemplatespecBuilder) WithNodeSelectorNew(nodeselectors map[string]string) *PodtemplatespecBuilder {
	if len(nodeselectors) == 0 {
		b.errs = append(
			b.errs,
			errors.New(
				"failed to build podtemplatespec object: missing nodeselectors",
			),
		)
		return b
	}

	// copy of original map
	newnodeselectors := map[string]string{}
	for key, value := range nodeselectors {
		newnodeselectors[key] = value
	}

	// override
	b.podtemplatespec.Object.Spec.NodeSelector = newnodeselectors
	return b
}

func (b *PodtemplatespecBuilder) WithNodeSelectorByValue(nodeselectors map[string]string) *PodtemplatespecBuilder {
	// copy of original map
	newnodeselectors := map[string]string{}
	for key, value := range nodeselectors {
		newnodeselectors[key] = value
	}

	// override
	b.podtemplatespec.Object.Spec.NodeSelector = newnodeselectors
	return b
}

// WithServiceAccountName sets the ServiceAccountnNme field of podtemplatespec
func (b *PodtemplatespecBuilder) WithServiceAccountName(serviceAccountnNme string) *PodtemplatespecBuilder {
	if len(serviceAccountnNme) == 0 {
		b.errs = append(
			b.errs,
			errors.New(
				"failed to build podtemplatespec object: missing serviceaccountname",
			),
		)
		return b
	}

	b.podtemplatespec.Object.Spec.ServiceAccountName = serviceAccountnNme
	return b
}

// WithAffinity sets the affinity field of podtemplatespec
func (b *PodtemplatespecBuilder) WithAffinity(affinity *corev1.Affinity) *PodtemplatespecBuilder {
	if affinity == nil {
		b.errs = append(
			b.errs,
			errors.New(
				"failed to build podtemplatespec object: missing affinity",
			),
		)
		return b
	}

	// copy of original pointer
	newaffinitylist := *affinity

	b.podtemplatespec.Object.Spec.Affinity = &newaffinitylist
	return b
}

// WithTolerationsByValue sets pod toleration.
// If provided tolerations argument is empty it does not complain.
func (b *PodtemplatespecBuilder) WithTolerationsByValue(tolerations ...corev1.Toleration) *PodtemplatespecBuilder {
	if len(b.podtemplatespec.Object.Spec.Tolerations) == 0 {
		b.podtemplatespec.Object.Spec.Tolerations = tolerations
		return b
	}

	b.podtemplatespec.Object.Spec.Tolerations = append(
		b.podtemplatespec.Object.Spec.Tolerations,
		tolerations...,
	)

	return b
}

// WithTolerations merges the existing tolerations
// with the provided arguments
func (b *PodtemplatespecBuilder) WithTolerations(tolerations ...corev1.Toleration) *PodtemplatespecBuilder {
	if tolerations == nil {
		b.errs = append(
			b.errs,
			errors.New(
				"failed to build podtemplatespec object: nil tolerations",
			),
		)
		return b
	}
	if len(tolerations) == 0 {
		b.errs = append(
			b.errs,
			errors.New(
				"failed to build podtemplatespec object: missing tolerations",
			),
		)
		return b
	}

	if len(b.podtemplatespec.Object.Spec.Tolerations) == 0 {
		return b.WithTolerationsNew(tolerations...)
	}

	b.podtemplatespec.Object.Spec.Tolerations = append(
		b.podtemplatespec.Object.Spec.Tolerations,
		tolerations...,
	)

	return b
}

// WithTolerationsNew sets the tolerations field of podtemplatespec
func (b *PodtemplatespecBuilder) WithTolerationsNew(tolerations ...corev1.Toleration) *PodtemplatespecBuilder {
	if tolerations == nil {
		b.errs = append(
			b.errs,
			errors.New(
				"failed to build podtemplatespec object: nil tolerations",
			),
		)
		return b
	}
	if len(tolerations) == 0 {
		b.errs = append(
			b.errs,
			errors.New(
				"failed to build podtemplatespec object: missing tolerations",
			),
		)
		return b
	}

	// copy of original slice
	newtolerations := []corev1.Toleration{}
	newtolerations = append(newtolerations, tolerations...)

	b.podtemplatespec.Object.Spec.Tolerations = newtolerations

	return b
}

// WithContainerBuilders builds the list of containerbuilder
// provided and merges it to the containers field of the podtemplatespec
func (b *PodtemplatespecBuilder) WithContainerBuilders(
	containerBuilderList ...*ContainerBuilder,
) *PodtemplatespecBuilder {
	if containerBuilderList == nil {
		b.errs = append(
			b.errs,
			errors.New("failed to build podtemplatespec: nil containerbuilder"),
		)
		return b
	}
	for _, containerBuilder := range containerBuilderList {
		containerObj, err := containerBuilder.Build()
		if err != nil {
			b.errs = append(
				b.errs,
				errors.Wrap(
					err,
					"failed to build podtemplatespec",
				),
			)
			return b
		}
		b.podtemplatespec.Object.Spec.Containers = append(
			b.podtemplatespec.Object.Spec.Containers,
			containerObj,
		)
	}
	return b
}

// WithVolumeBuilders builds the list of volumebuilders provided
// and merges it to the volumes field of podtemplatespec.
func (b *PodtemplatespecBuilder) WithVolumeBuilders(
	volumeBuilderList ...*VolumeBuilder,
) *PodtemplatespecBuilder {
	if volumeBuilderList == nil {
		b.errs = append(
			b.errs,
			errors.New("failed to build podtemplatespec: nil volumeBuilderList"),
		)
		return b
	}
	for _, volumeBuilder := range volumeBuilderList {
		vol, err := volumeBuilder.Build()
		if err != nil {
			b.errs = append(
				b.errs,
				errors.Wrap(err, "failed to build podtemplatespec"),
			)
			return b
		}
		newvol := *vol
		b.podtemplatespec.Object.Spec.Volumes = append(
			b.podtemplatespec.Object.Spec.Volumes,
			newvol,
		)
	}
	return b
}

// WithContainerBuildersNew builds the list of containerbuilder
// provided and sets the containers field of the podtemplatespec
func (b *PodtemplatespecBuilder) WithContainerBuildersNew(
	containerBuilderList ...*ContainerBuilder,
) *PodtemplatespecBuilder {
	if containerBuilderList == nil {
		b.errs = append(
			b.errs,
			errors.New("failed to build podtemplatespec: nil containerbuilder"),
		)
		return b
	}
	if len(containerBuilderList) == 0 {
		b.errs = append(
			b.errs,
			errors.New("failed to build podtemplatespec: missing containerbuilder"),
		)
		return b
	}
	containerList := []corev1.Container{}
	for _, containerBuilder := range containerBuilderList {
		containerObj, err := containerBuilder.Build()
		if err != nil {
			b.errs = append(
				b.errs,
				errors.Wrap(
					err,
					"failed to build podtemplatespec",
				),
			)
			return b
		}
		containerList = append(
			containerList,
			containerObj,
		)
	}
	b.podtemplatespec.Object.Spec.Containers = containerList
	return b
}

// WithVolumeBuildersNew builds the list of volumebuilders provided
// and sets Volumes field of podtemplatespec.
func (b *PodtemplatespecBuilder) WithVolumeBuildersNew(
	volumeBuilderList ...*VolumeBuilder,
) *PodtemplatespecBuilder {
	if volumeBuilderList == nil {
		b.errs = append(
			b.errs,
			errors.New("failed to build podtemplatespec: nil volumeBuilderList"),
		)
		return b
	}
	if len(volumeBuilderList) == 0 {
		b.errs = append(
			b.errs,
			errors.New("failed to build podtemplatespec: missing volumeBuilderList"),
		)
		return b
	}
	volList := []corev1.Volume{}
	for _, volumeBuilder := range volumeBuilderList {
		vol, err := volumeBuilder.Build()
		if err != nil {
			b.errs = append(
				b.errs,
				errors.Wrap(err, "failed to build podtemplatespec"),
			)
			return b
		}
		newvol := *vol
		volList = append(
			volList,
			newvol,
		)
	}
	b.podtemplatespec.Object.Spec.Volumes = volList
	return b
}

// WithRestartPolicy sets the RestartPolicy field of PodTemplateSpec
func (b *PodtemplatespecBuilder) WithRestartPolicy(restartPolicy corev1.RestartPolicy) *PodtemplatespecBuilder {
	if restartPolicy == "" {
		b.errs = append(
			b.errs,
			errors.New("failed to build podtemplatespec object: missing restart policy"),
		)
		return b
	}
	b.podtemplatespec.Object.Spec.RestartPolicy = restartPolicy
	return b
}


// Build returns a padTemplateSpec instance
func (b *PodtemplatespecBuilder) Build() (*PodTemplateSpec, error) {
	err := b.validate()
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"failed to build a podtemplatespec: %s",
			b.podtemplatespec.Object,
		)
	}
	return b.podtemplatespec, nil
}

func (b *PodtemplatespecBuilder) validate() error {
	if len(b.errs) != 0 {
		return errors.Errorf(
			"failed to validate: build errors were found: %v",
			b.errs,
		)
	}
	return nil
}
