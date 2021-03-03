# dev

Issue: https://github.com/amplify-edge/shared/issues/13

## functional needs we know

- devs need to trans their flutter arb files
- devs need to trans their golang code. prob later though.
- flutter app gui needs to be smaller and load the translations at runtime from the backend. At the moment they are embedded into the frontend.
  - so we need both because sys will want to embed, and module will want to not embed in general.
- flutter and golang need to trans their data, like what we use for bootstrapper and at runtime for Ion, and many other modules.
  - so server needs to provide a trans service.
- product and docs web site needs to be in all languages too
  - this presumes markdown.
- listmonk for outreach. https://github.com/knadh/listmonk
  - we need this for Verticals so Leyla can communicate with users in any language
  - we need this later in flutter, so GCN Orgs can do outreach. At this stage we will refactor GUI to use flutter.
  - It has a GUi that allows you to update the JSON that the web gui uses. https://github.com/knadh/listmonk/blob/master/i18n/en.json
    - so will need a parser for that eventualyl which is easy. Its just name value with ":" and CR delimting. EX: "admin.errorMarshallingConfig": "Error marshalling config: {error}",
  - The messages like emails need to be translated.
    - so will need it for that. And we will add more message channels over time.
- support
  - we need our video conf and chat to be multi-language to do support.

## design

- bs-lang we already have so match the API, so alex can get going.
  - func translate(text, from, to string, withVerification bool, tries int, delay time.Duration) (string, error) {
  - but make a NEW bs-lang as there is so much shit in it.  

- API
  - low level text to text, so that any tools above can do what they need to do and just call text to text as needed
  - this also means APi will not need to change.
  - GRPC makes sense for all this using our standard approach that we use for all services.
    - so add the GRPC APi to sys.shared, and allow shared/bs-lang to compile agaisnt it, as well as main.

- service for trans behind API.
  - catchzeng ( currently used by bs-lang )
  - goog
  - so because APi is simple, easy to match on goog and catchzeng

- storage and caching tiers
  - the service layer will use these storage mechanisms, so we lower costs. Essentially its acting as what they call "Translation Memory"
  - First call is to local store. If cache miss, go to Second call.
    - local genji store, so we can use it in the Server at runtime Or in CLI at dev time. memory or badger.
  - Second call is to Google Sheets as a "global HA store".  If cache miss, go to Third call.
  - Third call is to Google Trans itself.
  - so regarding Storage API - Genji and Google Sheets are a row/column store so easy to match the API.

- providence 
  - stored with data so that a bad translation can be traced backwards to the source.
  - make sure our API stores and returns the source of the translation, so bad translations can be fixed easily and so the translation memory gets smarter over time. This is why gsheets are nice as there is a GUI for anyone to fix things, and not just devs.
  - Might even make it visible in flutter later so make it easy to go backwards. Would probably be driven by a setting or feature flag.

- config
  - GOOGLE_CLOUD_PROJECT: see: https://github.com/googleapis/google-api-go-client/blob/master/internal/examples/mock/highlevel.go#L44


## Versioning

**stage 1**
bs-lang needs to work right now.
- do all the code in shared/tool/bs-lang-srv. Later will split the Service to sys-share, and make a GRPC API.
- just do the google sheets store and google translate API for now. That gets it working and not costing us a fortune.
- can all run locally, but use a global gsheets. The google service key can just be an env variable that we share with ech other for now.


## lib

https://github.com/s0nerik/goloc

cloud.google.com/go/translate
- https://github.com/search?l=Go&o=desc&q=cloud.google.com%2Fgo%2Ftranslate&s=indexed&type=Code


https://github.com/takakd/translation-api
- GUI: https://github.com/takakd/retranslate
  - typescript, but easy to make flutter, so dev can run as flutter app or app can use it.
- DEMO: https://retranslate-demo.herokuapp.com/
- grpc api
- aws and google
- deployments local and cloud. uses envoy.


**sheets**

https://developers.google.com/sheets/api/quickstart/go
go get -u google.golang.org/api/sheets/v4
go get -u golang.org/x/oauth2/google

https://github.com/rr250/jobpost-plugin/blob/master/server/google.go
- drive and sheets API. Need both.
- designed to maintain mailing lsit subscriptions.





## other

Flutter: https://github.com/aloisdeniel/flutter_sheet_localization

https://github.com/bratan/flutter_translate
- uses https://github.com/bratan/flutter_device_locale
https://github.com/bratan/flutter_translate_gen

Golang Gsheets: https://github.com/Iwark/spreadsheet/blob/v2/service.go




## Ops - Cloud run

Set it up to deploy on git commit.
https://cloud.google.com/run/docs/continuous-deployment-with-cloud-build



## Ops

ToDO

proper logging
"go.uber.org/zap"
https://github.com/knative/serving/blob/master/pkg/reconciler/autoscaling/hpa/hpa.go
- uses knative.dev/pkg/logging
https://github.com/knative/pkg/tree/master/metrics
https://github.com/knative/pkg/tree/master/logging


propper logging, tracing and metrics
https://github.com/jaegertracing/jaeger/tree/master/examples/hotrod
- this looks strong
- Here is how to setup the entry point: https://github.com/jaegertracing/jaeger/blob/master/examples/hotrod/cmd/root.go#L77