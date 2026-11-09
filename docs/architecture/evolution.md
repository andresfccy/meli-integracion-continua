# Architecture Evolution

This document separates approved current architecture from future options. Future options are not implemented behavior unless a spec, decision, plan, and code change make them current.

## Current implementation

- Go monolith backend with lightweight hexagonal boundaries.
- Vite React TypeScript frontend with feature-based modules.
- In-memory persistence for the MVP.
- REST commands and authenticated read-only GraphQL composition.
- Seeded authentication for Seller, Buyer, and Operations roles.
- Operations priority queue and manual/scheduled priority review.
- Driver-backed console delivery for seller notification events.

## Approved direction

- Keep boundaries explicit and simple.
- Keep business rules in domain/application layers.
- Keep HTTP, GraphQL, scheduler, storage, and notification protocols at adapter boundaries.
- Keep specs and decisions as the basis for future behavior changes.
- Prefer deterministic behavior and small, independently verifiable changes.

## Future options

These options are intentionally not current commitments:

- Durable persistence.
- External notification providers with retries and delivery tracking.
- Buyer-facing notifications.
- Pagination for larger queues and order lists.
- Public webhooks or external buyer-question ingestion.
- GraphQL mutations, subscriptions, federation, or public GraphQL contracts.
- gRPC service boundaries if the monolith is split.
- Event streaming and read-optimized storage.
- Infrastructure as Code and production deployment topology.

## Evolution rule

Before implementing a future option, update or add the relevant spec, decision, implementation plan entry, and verification criteria. Do not treat this document as authorization to implement future behavior by itself.
