package ServerCore

import (
	"net/http"
	"testing"

	"TranslateServer/internal/ServerPlatform/mock"
	"TranslateServer/internal/Supervisor/api"
	"TranslateServer/internal/Supervisor/mock"
)

func TestMetricHandler_Handle_ReturnsMetrics(t *testing.T) {
	mockSupervisor := new(MockSupervisorApi.MockSupervisor)
	metricHandler := &MetricHandler{mockSupervisor}
	mockHandler := new(MockServerInterface.MockServerInterface)

	metrics := []SupervisorApi.ChannelStatus{
		{Id: "m1", Status: SupervisorApi.WORKING},
		{Id: "m2", Status: SupervisorApi.NOT_FOUND},
		{Id: "m3", Status: SupervisorApi.DEFECT},
		{Id: "m4", Status: 999},
	}

	expectedResult := map[string]string{
		"m1": "WORKING",
		"m2": "NOT_FOUND",
		"m3": "DEFECT",
		"m4": "UNKNOWN",
	}

	mockSupervisor.On("GetMetric").Return(metrics)
	mockHandler.On("JsonCallback", http.StatusOK, expectedResult).Return()

	metricHandler.Handle(mockHandler)

	mockSupervisor.AssertExpectations(t)
	mockHandler.AssertExpectations(t)
}

func TestMetricHandler_Handle_EmptyMetrics(t *testing.T) {
	mockSupervisor := new(MockSupervisorApi.MockSupervisor)
	metricHandler := &MetricHandler{mockSupervisor}
	mockHandler := new(MockServerInterface.MockServerInterface)

	expectedResult := map[string]string{}

	mockSupervisor.On("GetMetric").Return([]SupervisorApi.ChannelStatus{})
	mockHandler.On("JsonCallback", http.StatusOK, expectedResult).Return()

	metricHandler.Handle(mockHandler)

	mockSupervisor.AssertExpectations(t)
	mockHandler.AssertExpectations(t)
}
