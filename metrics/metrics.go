package metrics

import (
    "github.com/gofiber/fiber/v2"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "net/http"
)

var (
    reqsCounter = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Número total de requisições",
        },
        []string{"path"},
    )
    latencyHistogram = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name: "http_request_duration_seconds",
            Help: "Duração das requisições HTTP",
        },
        []string{"path"},
    )
)

func SetupPrometheus(app *fiber.App) {
    prometheus.MustRegister(reqsCounter, latencyHistogram)

    app.Use(func(c *fiber.Ctx) error {
        timer := prometheus.NewTimer(latencyHistogram.WithLabelValues(c.Path()))
        defer timer.ObserveDuration()

        err := c.Next()
        reqsCounter.WithLabelValues(c.Path()).Inc()
        return err
    })

    app.Get("/metrics", fiber.WrapH(promhttp.Handler()))
}
