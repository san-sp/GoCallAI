# GoCallAI architecture

## Core principle

The telephony provider and cloud backend answer forwarded calls. The Android app is a control plane, not part of the critical audio path, and does not need to remain open.

## Call path

1. Caller dials the user's existing mobile number.
2. Carrier-side conditional forwarding sends unanswered/busy/unreachable calls to a provisioned telephony number.
3. The telephony provider invokes the configured inbound call flow.
4. The provider establishes a bidirectional media stream to the backend voice gateway, subject to provider support and account enablement.
5. The backend creates a call session and connects the selected realtime voice engine.
6. At call completion, the backend persists permitted metadata, transcript, and summary, then sends a push notification.
7. If recording is explicitly enabled, a media pipeline stores caller and assistant audio in private object storage. Recording stays off by default until disclosure, consent, provider retention, and applicable legal requirements are reviewed.

## Control plane and call plane

- Control plane: Android app -> authenticated Go API -> settings and mode requests.
- Call plane: carrier -> telephony provider -> backend media gateway -> AI engine.
- Requested assistant mode and verified carrier forwarding state are separate. Never display forwarding as active merely because a database flag changed.

## Initial modes

- `automatic`: carrier-supported conditional forwarding rules.
- `manual_active`: temporarily forward all calls only after the user completes the supported carrier activation flow and the backend verifies it.
- `disabled`: assistant is not intended to receive calls.

Carrier forwarding may require the Android dialer, carrier app, or carrier-supported MMI workflow. Do not assume an Android app can silently change network forwarding settings.

## Security boundaries

- Verify provider webhook authentication before acting on callbacks.
- Treat provider call IDs as idempotency keys.
- Keep provider and AI credentials out of source control and Android resources.
- Authorize call recordings, transcripts, and playback to the owning user.
- Use short-lived signed playback URLs or an authenticated streaming endpoint.
- Encrypt recordings at rest and in transit; define retention and deletion.
- Do not record by default. Provide audible disclosure and obtain required consent before recording.
