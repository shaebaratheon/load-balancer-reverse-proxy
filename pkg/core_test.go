package pkg

import (
	"testing"
	"time"
)

func TestProxyCoreLifecycle(t *testing.T) {
	cfg := ProxyCoreConfig{TimeoutMs: 1000, MaxRetries: 3, Enabled: true}
	eng := NewProxyCoreEngine(cfg)
	ctx := &ProxyCoreContext{
		CorrelationID: "cid-proxy_core-01",
		Payload: map[string]string{"k1": "v1"},
		Timestamp: time.Now(),
	}
	digest, err := eng.Process(ctx)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if len(digest) == 0 {
		t.Fatal("expected valid digest")
	}
}

func TestBalancerAlgorithmsLifecycle(t *testing.T) {
	cfg := BalancerAlgorithmsConfig{TimeoutMs: 1000, MaxRetries: 3, Enabled: true}
	eng := NewBalancerAlgorithmsEngine(cfg)
	ctx := &BalancerAlgorithmsContext{
		CorrelationID: "cid-balancer_algorithms-01",
		Payload: map[string]string{"k1": "v1"},
		Timestamp: time.Now(),
	}
	digest, err := eng.Process(ctx)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if len(digest) == 0 {
		t.Fatal("expected valid digest")
	}
}

func TestHealthCheckerLifecycle(t *testing.T) {
	cfg := HealthCheckerConfig{TimeoutMs: 1000, MaxRetries: 3, Enabled: true}
	eng := NewHealthCheckerEngine(cfg)
	ctx := &HealthCheckerContext{
		CorrelationID: "cid-health_checker-01",
		Payload: map[string]string{"k1": "v1"},
		Timestamp: time.Now(),
	}
	digest, err := eng.Process(ctx)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if len(digest) == 0 {
		t.Fatal("expected valid digest")
	}
}

func TestTlsTerminationLifecycle(t *testing.T) {
	cfg := TlsTerminationConfig{TimeoutMs: 1000, MaxRetries: 3, Enabled: true}
	eng := NewTlsTerminationEngine(cfg)
	ctx := &TlsTerminationContext{
		CorrelationID: "cid-tls_termination-01",
		Payload: map[string]string{"k1": "v1"},
		Timestamp: time.Now(),
	}
	digest, err := eng.Process(ctx)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if len(digest) == 0 {
		t.Fatal("expected valid digest")
	}
}

func TestCircuitBreakerLifecycle(t *testing.T) {
	cfg := CircuitBreakerConfig{TimeoutMs: 1000, MaxRetries: 3, Enabled: true}
	eng := NewCircuitBreakerEngine(cfg)
	ctx := &CircuitBreakerContext{
		CorrelationID: "cid-circuit_breaker-01",
		Payload: map[string]string{"k1": "v1"},
		Timestamp: time.Now(),
	}
	digest, err := eng.Process(ctx)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if len(digest) == 0 {
		t.Fatal("expected valid digest")
	}
}

func TestAccessLoggerLifecycle(t *testing.T) {
	cfg := AccessLoggerConfig{TimeoutMs: 1000, MaxRetries: 3, Enabled: true}
	eng := NewAccessLoggerEngine(cfg)
	ctx := &AccessLoggerContext{
		CorrelationID: "cid-access_logger-01",
		Payload: map[string]string{"k1": "v1"},
		Timestamp: time.Now(),
	}
	digest, err := eng.Process(ctx)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if len(digest) == 0 {
		t.Fatal("expected valid digest")
	}
}

