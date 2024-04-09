package request

import (
	"mygin/models/base"
)

type Base struct {
	Name string `json:"name"`
	Labels []base.ListMapItem `json:"labels"`
	Namespace string `json:"namespace"`
	RestartPolicy string `json:"restartPolicy"`
}


type Pod struct {
	Base           Base                `json:"base"`
	Tolerations    []corev1.Toleration `json:"toleartions"`
	NodeScheduling Nodescheduling      `json:"nodeScheduling"`
	Volume         []Volume            `json:"volumes"`
	NetWorking     NetWorking          `json:"netWorking"`
	InitContainers []Container         `json:"initContainers"`
	Containers     []Container         `json:"container"`
}
