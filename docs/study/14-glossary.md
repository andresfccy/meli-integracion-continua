# Glossary

| Term | Plain explanation | Meaning here |
| --- | --- | --- |
| Domain | Core rules that remain if screens/providers change | orders, questions, priority |
| Application | Coordinator for a user/system goal | login, create, reply, approve, review |
| Use case | One intentional system operation | approve an answer |
| Port | Interface describing a needed capability | `Store`, `Delivery` |
| Adapter | Code translating a protocol/technology | HTTP router, memory store |
| Driver | Outbound adapter for a provider/channel | console email delivery |
| Dependency | Code one component needs | application depends on domain |
| Dependency inversion | Inner policy defines what it needs; outer implementation plugs in | application interface, memory implementation |
| Monolith | One deployable process | Go API with internal modules |
| Microservice | Independently deployed network service | deliberately deferred |
| REST | HTTP resource/command style | state-changing endpoints |
| GraphQL | Schema/query language for selected data | schema-backed, read-only role composition with no mutations |
| Command | Request to change state with intent | reply, approve, review |
| Query | Request to read without changing state | dashboard/order detail |
| JWT | Signed compact claims token | one-hour HS256 token validated for header, signature, claims, and seeded subject |
| CORS | Browser rule controlling cross-origin calls | allow port 3000/preview to call 8080 |
| Persistence | Keeping data | current memory lasts until restart |
| Determinism | Same inputs produce same result/order | priority and queue sorting |
| Idempotency | Repetition does not duplicate an effect | atomic one-execution-per-UTC-window scheduling |
| MVP | Small version validating essential behavior | seeds, memory, console delivery |
| ADR | Record of context, decision, alternatives, consequences | architecture decision files |
| Traceability | Link from requirement to evidence | specs → code → tests |
| Technical debt | Current shortcut that raises future cost/risk | browser-test gap, seed security, in-memory infrastructure |
| Inference | Reasonable conclusion not explicitly recorded | React rationale |
| Evidence | Source supporting a claim | ADR, code, passing test |
| Runtime validation | Checking external data while software runs | missing for frontend API payloads |
| Seed data | Fixed demonstration data | users, orders, question, channels |
| Coarse status | Broad compatibility state | open/resolved |
| Lifecycle status | Detailed workflow phase | question_active/seller_answered/closed |
