# load-balancer-reverse-proxy

Layer 7 HTTP/gRPC reverse proxy with round-robin, least-connections, and active health checks in Go.

## Architecture & Design

This project implements a high-reliability distributed architecture designed for production workloads.
### Core Components
- `proxy_core`: Core subsystem handling specific domain logic, invariants, and performance guarantees.
- `balancer_algorithms`: Core subsystem handling specific domain logic, invariants, and performance guarantees.
- `health_checker`: Core subsystem handling specific domain logic, invariants, and performance guarantees.
- `tls_termination`: Core subsystem handling specific domain logic, invariants, and performance guarantees.
- `circuit_breaker`: Core subsystem handling specific domain logic, invariants, and performance guarantees.
- `access_logger`: Core subsystem handling specific domain logic, invariants, and performance guarantees.

## Testing and Verification

Run the test suite via standard tooling.
