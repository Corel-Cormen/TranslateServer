package ServerCore

import (
	"net/http"

	"TranslateServer/internal/ServerPlatform/api"
	"TranslateServer/internal/Supervisor/api"
)

type MetricHandler struct {
	supervisorInterface SupervisorApi.SupervisorInterface
}

func metricStatusStr(status int) string {
	switch status {
	case SupervisorApi.WORKING:
		return "WORKING"
	case SupervisorApi.NOT_FOUND:
		return "NOT_FOUND"
	case SupervisorApi.DEFECT:
		return "DEFECT"
	default:
		return "UNKNOWN"
	}
}

func (h *MetricHandler) Handle(handler ServerCoreApi.HandlerInterface) {
	metrics := h.supervisorInterface.GetMetric()
	result := make(map[string]string, len(metrics))
	for _, metric := range metrics {
		result[metric.Id] = metricStatusStr(metric.Status)
	}
	handler.JsonCallback(http.StatusOK, result)
}
