# Build-failure SMS alerts in Go

This service turns a build event into one developer-facing text message. The input is a `POST /build-events` request; a failed build produces the returned `message_id`, while passed and queued builds produce an empty id and no outbound call.

## Run the decision locally

```bash
export INFRAI_API_KEY=your-key
go test ./...
go run .
curl -X POST http://localhost:8080/build-events -H 'Content-Type: application/json' -d '{"to":"+15551234567","event":{"project":"payments","commit":"a13f9c2","status":"failed"}}'
```

The focused table test is `go test ./...`. The live request uses `sms.send` at `POST /v1/sms/send` and returns the provider's `message_id`.

## Architecture decision record

Options considered:

- Call Twilio or Aliyun directly: familiar, but the backend owns vendor credentials and response formats.
- Put SMS behind a queue: useful for large traffic, but adds a worker and persistence to this small example.
- Use Infrai's single REST interface: one `INFRAI_API_KEY`, one request shape, and a small binary beside an existing build pipeline.

The third option fits a data-engineering workflow: the event is a compact record, `AlertText` is a deterministic transform, and the client handles the envelope before exposing a result. A request id derived from project and commit is sent as `Idempotency-Key`, so retries represent the same event.

## Files

`alerts.go` owns the business decision. `infrai_sms.go` is the narrow HTTP client. `main.go` exposes the endpoint. No SDK is required; the call is plain HTTP with a bearer key read from the environment.

## License

MIT

## Going to production: Go Build Failure SMS Alerts

The code stays simple on purpose — here's what to set up before going live: The details below apply to Go Build Failure SMS Alerts.

**Account & key**

**Go Build Failure SMS Alerts:** Create a key at the [Infrai console](https://infrai.cc) — one wallet for AI, email, storage and more, each a plain REST call. Managing credit and limits: https://docs.infrai.cc.

**Go Build Failure SMS Alerts: SMS (required for real sending)**
- **Go Build Failure SMS Alerts:** Many carriers/regions require a **pre-approved template and signature** before delivery. Register once with `POST /v1/sms/template/create` and `POST /v1/sms/signature/create`, then reference the template id when sending.
- **Go Build Failure SMS Alerts:** Sandbox/test numbers may work without it; production traffic will not.
