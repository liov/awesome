module test

go 1.22

toolchain go1.22.0

require (
	bou.ke/monkey v1.0.2
	github.com/brahma-adshonor/gohook v1.1.9
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc
	github.com/google/wire v0.6.0
	github.com/hopeio/cherry v1.0.0
	github.com/lucasb-eyer/go-colorful v1.0.2
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.18.2
	github.com/unicorn-engine/unicorn v0.0.0-20191119163456-3cea38bff7bf
	github.com/xuri/excelize/v2 v2.8.1
	go.uber.org/fx v1.21.0
	go.uber.org/zap v1.27.0
	golang.org/x/exp v0.0.0-20240416160154-fe59bbe5cc7f
	golang.org/x/exp/shiny v0.0.0-20220414153411-bcd21879b8fd
	golang.org/x/mobile v0.0.0-20201217150744-e6ae53a27f4f
	golang.org/x/sys v0.19.0
	gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v3 v3.0.1
	gorm.io/gorm v1.25.10
)

require (
	dmitri.shuralyov.com/gpu/mtl v0.0.0-20201218220906-28db891af037 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20200222043503-6f7a984d4dc4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20231201235250-de7065d80cb9 // indirect
	github.com/jackc/pgx/v5 v5.5.5 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jezek/xgb v1.0.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/pelletier/go-toml/v2 v2.2.1 // indirect
	github.com/prometheus/client_golang v1.19.0 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.53.0 // indirect
	github.com/prometheus/procfs v0.14.0 // indirect
	github.com/richardlehane/mscfb v1.0.4 // indirect
	github.com/richardlehane/msoleps v1.0.3 // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/xuri/efp v0.0.0-20240408161823-9ad904a10d6d // indirect
	github.com/xuri/nfp v0.0.0-20240318013403-ab9948c2c4a7 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/dig v1.17.1 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/arch v0.7.0 // indirect
	golang.org/x/crypto v0.22.0 // indirect
	golang.org/x/image v0.15.0 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gorm.io/driver/postgres v1.5.7 // indirect
	gorm.io/plugin/prometheus v0.1.0 // indirect
)

replace (
	github.com/hopeio/cherry => ../../../thirdparty/cherry
	gorm.io/gorm => ../../../../../gorm
)
