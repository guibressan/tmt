package tmt

var Tech = map[string]string{
	"aws":               "AWS",
	"aws_lambda":        "AWS_Lambda",
	"actionscript":      "ActionScript",
	"aerospike":         "Aerospike",
	"amazon_cloudfront": "Amazon_CloudFront",
	"amazon_ec2":        "Amazon_EC2",
	"amazon_ecs":        "Amazon_ECS",
	"amazon_kinesis":    "Amazon_Kinesis",
	"angular":           "Angular",
	"angulardart":       "AngularDart",
	"apachedoris":       "ApacheDoris",
	"apachedruid":       "ApacheDruid",
	"apache_kafka":      "Apache_Kafka",
	"astro":             "Astro",
	"aurora":            "Aurora",
	"aurora_postgresql": "Aurora_PostgreSQL",
	"bigquery":          "BigQuery",
	"c":                 "C",
	"c#":                "C#",
	"cpp":               "CPP",
	"caddy":             "Caddy",
	"cassandra":         "Cassandra",
	"citusdb":           "CitusDB",
	"clickhouse":        "Clickhouse",
	"clickhousekeeper":  "ClickhouseKeeper",
	"clojure":           "Clojure",
	"cloud":             "Cloud",
	"cloudflare":        "CloudFlare",
	"cloudsql":          "CloudSQL",
	"cockroachdb":       "CockroachDB",
	"coffeescript":      "CoffeeScript",
	"commonlisp":        "CommonLISP",
	"consul":            "Consul",
	"couchbase":         "Couchbase",
	"db2":               "DB2",
	"dart":              "Dart",
	"debian":            "Debian",
	"dockerswarm":       "DockerSwarm",
	"dotnet":            "DotNET",
	"dotnetcore":        "DotNetCore",
	"druid":             "Druid",
	"dynamodb":          "DynamoDB",
	"elasticsearch":     "ElasticSearch",
	"elasticache_redis": "Elasticache_Redis",
	"electron":          "Electron",
	"elixir":            "Elixir",
	"elm":               "Elm",
	"elsticsearch":      "ElsticSearch",
	"emberjs":           "EmberJS",
	"encore":            "Encore",
	"envoy":             "Envoy",
	"erlang":            "Erlang",
	"express":           "Express",
	"firebase":          "Firebase",
	"firecracker":       "Firecracker",
	"flex":              "Flex",
	"flutter":           "Flutter",
	"gwt":               "GWT",
	"gatsby":            "Gatsby",
	"go":                "Go",
	"graphql":           "GraphQL",
	"haproxy":           "HAProxy",
	"hbase":             "HBase",
	"htmx":              "HTMX",
	"hack":              "Hack",
	"hadoop":            "Hadoop",
	"hasura":            "Hasura",
	"heroku":            "Heroku",
	"hetzner":           "Hetzner",
	"hive":              "Hive",
	"hugo":              "Hugo",
	"influxdb":          "InfluxDB",
	"jruby":             "JRuby",
	"jsdoc":             "JSDoc",
	"json":              "JSON",
	"java":              "Java",
	"javascript":        "JavaScript",
	"kafka":             "Kafka",
	"kibana":            "Kibana",
	"kotlin":            "Kotlin",
	"kubernetes":        "Kubernetes",
	"kuma":              "Kuma",
	"ledgerstore":       "LedgerStore",
	"linux":             "Linux",
	"logstash":          "LogStash",
	"loki":              "Loki",
	"mssqlserver":       "MSSQLServer",
	"mac":               "Mac",
	"mariadb":           "MariaDB",
	"meteor":            "Meteor",
	"mongodb":           "MongoDB",
	"mysql":             "MySQL",
	"nats":              "NATS",
	"native":            "Native",
	"neo4j":             "Neo4j",
	"nginx":             "Nginx",
	"nhost":             "Nhost",
	"nodejs":            "NodeJS",
	"nomad":             "Nomad",
	"nuraft":            "NuRaft",
	"ovh":               "OVH",
	"objc":              "ObjC",
	"oracle":            "Oracle",
	"oracleexadata":     "OracleExadata",
	"oracle_timesten":   "Oracle_TimesTen",
	"php":               "PHP",
	"perl":              "Perl",
	"polymer":           "Polymer",
	"pony":              "Pony",
	"postgresql":        "PostgreSQL",
	"preact":            "Preact",
	"presto":            "Presto",
	"prisma":            "Prisma",
	"prometheus":        "Prometheus",
	"protobuf":          "ProtoBuf",
	"python":            "Python",
	"rabbitmq":          "RabbitMQ",
	"raft":              "Raft",
	"react":             "React",
	"reactnative":       "ReactNative",
	"redhat":            "RedHat",
	"redpanda":          "RedPanda",
	"redis":             "Redis",
	"redshift":          "Redshift",
	"render":            "Render",
	"riak":              "Riak",
	"rockset":           "Rockset",
	"rockylinux":        "RockyLinux",
	"rollup":            "RollUp",
	"ruby":              "Ruby",
	"rubyonrails":       "RubyOnRails",
	"rust":              "Rust",
	"s3_metadata":       "S3_metadata",
	"scala":             "Scala",
	"scylladb":          "ScyllaDB",
	"self-managed":      "Self-Managed",
	"sidekiq":           "Sidekiq",
	"singlestore":       "SingleStore",
	"skew":              "Skew",
	"snowpack":          "SnowPack",
	"solr":              "Solr",
	"supabase":          "Supabase",
	"svelte":            "Svelte",
	"sveltekit":         "SvelteKit",
	"swift":             "Swift",
	"tarantool":         "Tarantool",
	"thanos":            "Thanos",
	"tidb":              "TiDB",
	"timescaledb":       "TimescaleDB",
	"typescript":        "TypeScript",
	"ubuntu_core":       "Ubuntu_Core",
	"unknown":           "Unknown",
	"victoriametrics":   "VictoriaMetrics",
	"vite":              "Vite",
	"vitess":            "Vitess",
	"wasm":              "WASM",
	"wails.io":          "Wails.io",
	"webpack":           "WebPack",
	"windows":           "Windows",
	"yugabytedb":        "YugabyteDB",
	"zeromq":            "ZeroMQ",
	"zig":               "Zig",
	"zookeeper":         "ZooKeeper",
	"flyd":              "flyd",
	"grpc":              "gRPC",
	"jquery":            "jQuery",
	"ksqldb":            "kSqlDB",
	"mrsk":              "mrsk",
	"nindex":            "nIndex",
	"trpc":              "tRPC",
}

var migrations = []migration{
	// from https://github.com/kokizzu/list-of-tech-migrations
	{
		Company: "Reddit",
		Uri:     "https://redditblog.com/2005/12/05/on-lisp/",
		Year:    2005,
		From: []string{
			"commonlisp",
		},
		To: []string{
			"python",
		},
	},
	{
		Company: "Bloomberg",
		Uri:     "https://www.bloomberg.com/company/stories/10-insights-adopting-typescript-at-scale/",
		Year:    2005,
		From: []string{
			"c",
			"cpp",
		},
		To: []string{
			"javascript",
		},
	},
	{
		Company: "Bing",
		Uri:     "https://devblogs.microsoft.com/dotnet/migration-of-bings-workflow-engine-to-net-5/",
		Year:    2010,
		From: []string{
			"cpp",
		},
		To: []string{
			"dotnet",
		},
	},
	{
		Company: "Twitter",
		Uri:     "http://readwrite.com/2011/07/06/twitter-java-scala/",
		Year:    2011,
		From: []string{
			"ruby",
		},
		To: []string{
			"scala",
		},
	},
	{
		Company: "UrbanAirship",
		Uri:     "http://wiki.postgresql.org/images/7/7f/Adam-lowry-postgresopen2011.pdf",
		Year:    2011,
		From: []string{
			"mongodb",
		},
		To: []string{
			"postgresql",
		},
	},
	{
		Company: "Mixpanel",
		Uri:     "https://code.mixpanel.com/2011/08/05/how-and-why-we-switched-from-erlang-to-python/",
		Year:    2011,
		From: []string{
			"erlang",
		},
		To: []string{
			"python",
		},
	},
	{
		Company: "LinkedIn",
		Uri:     "http://highscalability.com/blog/2012/10/4/linkedin-moved-from-rails-to-node-27-servers-cut-and-up-to-2.html",
		Year:    2012,
		From: []string{
			"ruby",
		},
		To: []string{
			"nodejs",
		},
	},
	{
		Company: "MoovWeb",
		Uri:     "https://groups.google.com/forum/#!topic/golang-nuts/MeiTNnGhLg8/discussion",
		Year:    2012,
		From: []string{
			"unknown",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "NTPPool",
		Uri:     "https://news.ntppool.org/2012/10/new-dns-server/",
		Year:    2012,
		From: []string{
			"perl",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "DLGoogle",
		Uri:     "https://talks.golang.org/2013/oscon-dl.slide#10",
		Year:    2012,
		From: []string{
			"cpp",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "Bitly",
		Uri:     "http://word.bitly.com/post/29550171827/go-go-gadget",
		Year:    2012,
		From: []string{
			"python",
			"c",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "SoundCloud",
		Uri:     "https://developers.soundcloud.com/blog/go-at-soundcloud",
		Year:    2012,
		From: []string{
			"ruby",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "SmartyStreets",
		Uri:     "https://blog.gopheracademy.com/birthday-bash-2014/building-street-address-autocomplete/",
		Year:    2012,
		From: []string{
			"dotnet",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "Toggl",
		Uri:     "https://blog.toggl.com/2012/09/moving-to-go/",
		Year:    2012,
		From: []string{
			"ruby",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "Shareaholic",
		Uri:     "http://www.slideshare.net/Shareaholic/migrating-to-riak-at-shareaholic",
		Year:    2012,
		From: []string{
			"mongodb",
		},
		To: []string{
			"riak",
		},
	},
	{
		Company: "DigiDoc",
		Uri:     "http://svs.io/post/31724990463/why-i-migrated-away-from-mongodb",
		Year:    2012,
		From: []string{
			"mongodb",
		},
		To: []string{
			"postgresql",
		},
	},
	{
		Company: "Etsy",
		Uri:     "http://mcfunley.com/why-mongodb-never-worked-out-at-etsy",
		Year:    2012,
		From: []string{
			"mongodb",
		},
		To: []string{
			"mysql",
		},
	},
	{
		Company: "TekPub",
		Uri:     "http://rob.conery.io/2012/02/22/alt-tekpub-moving-to-mongodb/",
		Year:    2012,
		From: []string{
			"mongodb",
		},
		To: []string{
			"postgresql",
		},
	},
	{
		Company: "TargeterApp",
		Uri:     "http://blog.targeterapp.com/post/22984987832/why-we-moved-from-nodejs-to-ror",
		Year:    2012,
		From: []string{
			"nodejs",
		},
		To: []string{
			"ruby",
		},
	},
	{
		Company: "Dropbox",
		Uri:     "https://dropbox.tech/application/dropbox-dives-into-coffeescript",
		Year:    2012,
		From: []string{
			"javascript",
		},
		To: []string{
			"coffeescript",
		},
	},
	{
		Company: "PayPal",
		Uri:     "https://www.paypal-engineering.com/2013/11/22/node-js-at-paypal/",
		Year:    2013,
		From: []string{
			"java",
		},
		To: []string{
			"nodejs",
		},
	},
	{
		Company: "Koding",
		Uri:     "https://www.quora.com/Why-did-Koding-switch-from-Node-js-to-Go",
		Year:    2013,
		From: []string{
			"nodejs",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "Iron.io",
		Uri:     "https://www.iron.io/how-we-went-from-30-servers-to-2-go/",
		Year:    2013,
		From: []string{
			"ruby",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "Zalora",
		Uri:     "http://www.slideshare.net/wuvist1/zalora-php-togoen",
		Year:    2013,
		From: []string{
			"php",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "TimeHop",
		Uri:     "https://medium.com/building-timehop/why-timehop-chose-go-to-replace-our-rails-app-2855ea1912d#.h9cc85kym",
		Year:    2013,
		From: []string{
			"ruby",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "CanonicalJuju",
		Uri:     "https://groups.google.com/forum/#!topic/golang-nuts/jLnMsUbYwrQ",
		Year:    2013,
		From: []string{
			"python",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "Facebook",
		Uri:     "https://en.wikipedia.org/wiki/HHVM",
		Year:    2013,
		From: []string{
			"php",
		},
		To: []string{
			"hack",
		},
	},
	{
		Company: "Blippex",
		Uri:     "https://blippex.github.io/updates/2013/09/23/why-we-moved-away-from-aws.html",
		Year:    2013,
		From: []string{
			"aws",
		},
		To: []string{
			"ovh",
		},
	},
	{
		Company: "Instill",
		Uri:     "https://www.youtube.com/watch?v=Hg3cEBgq9Ds",
		Year:    2014,
		From: []string{
			"angulardart",
		},
		To: []string{
			"polymer",
		},
	},
	{
		Company: "Workia",
		Uri:     "https://www.youtube.com/watch?v=4O4jr0tr_ow",
		Year:    2014,
		From: []string{
			"actionscript",
		},
		To: []string{
			"dart",
		},
	},
	{
		Company: "TrustWave",
		Uri:     "https://www.youtube.com/watch?v=5-32KP0JHaE",
		Year:    2014,
		From: []string{
			"actionscript",
			"flex",
		},
		To: []string{
			"dart",
		},
	},
	{
		Company: "Facebook",
		Uri:     "https://www.facebook.com/notes/facebook-engineering/chat-stability-and-scalability/51412338919/",
		Year:    2014,
		From: []string{
			"erlang",
		},
		To: []string{
			"cpp",
		},
	},
	{
		Company: "Grab",
		Uri:     "https://www.youtube.com/watch?v=L688sHqXL2A",
		Year:    2014,
		From: []string{
			"ruby",
			"nodejs",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "CloudFlare",
		Uri:     "https://blog.cloudflare.com/scaling-out-postgresql-for-cloudflare-analytics-using-citusdb/",
		Year:    2014,
		From: []string{
			"postgresql",
		},
		To: []string{
			"citusdb",
		},
	},
	{
		Company: "Auth0",
		Uri:     "https://tomasz.janczuk.org/2015/09/from-kafka-to-zeromq-for-log-aggregation.html",
		Year:    2015,
		From: []string{
			"apache_kafka",
		},
		To: []string{
			"zeromq",
		},
	},
	{
		Company: "Pinterest",
		Uri:     "https://venturebeat.com/2015/12/18/pinterest-elixir/",
		Year:    2015,
		From: []string{
			"java",
		},
		To: []string{
			"elixir",
		},
	},
	{
		Company: "PresidentUniversity",
		Uri:     "https://www.socketloop.com/blogs/interview-with-kiswono-prayogo-head-of-software-development-at-president-university-indonesia",
		Year:    2015,
		From: []string{
			"php",
			"mysql",
		},
		To: []string{
			"go",
			"postgresql",
		},
	},
	{
		Company: "UserLike",
		Uri:     "https://www.userlike.com/en/blog/bye-by-mysql-and-mongodb-guten-tag-postgresql",
		Year:    2015,
		From: []string{
			"mysql",
			"mongodb",
		},
		To: []string{
			"postgresql",
		},
	},
	{
		Company: "Olery",
		Uri:     "http://developer.olery.com/blog/goodbye-mongodb-hello-postgresql/",
		Year:    2015,
		From: []string{
			"mongodb",
		},
		To: []string{
			"postgresql",
		},
	},
	{
		Company: "CrowdStrike",
		Uri:     "https://web.archive.org/web/20160612120018/http://jimplush.com/talk/2015/12/19/moving-a-team-from-scala-to-golang/",
		Year:    2015,
		From: []string{
			"scala",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "DockYard",
		Uri:     "https://dockyard.com/blog/2015/11/18/phoenix-is-not-rails",
		Year:    2015,
		From: []string{
			"ruby",
		},
		To: []string{
			"elixir",
		},
	},
	{
		Company: "Parse",
		Uri:     "https://web.archive.org/web/20150611021959/http://blog.parse.com/learn/how-we-moved-our-api-from-ruby-to-go-and-saved-our-sanity/",
		Year:    2015,
		From: []string{
			"ruby",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "Marriot",
		Uri:     "https://diginomica.com/why-marriott-is-transforming-their-legacy-systems-with-nosql",
		Year:    2015,
		From: []string{
			"unknown",
		},
		To: []string{
			"couchbase",
		},
	},
	{
		Company: "CrowdStrike",
		Uri:     "http://126kr.com/article/8sx2b2nrcc7",
		Year:    2016,
		From: []string{
			"scala",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "Tokopedia",
		Uri:     "http://tech.tokopedia.com/blog/perl-to-go/",
		Year:    2016,
		From: []string{
			"perl",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "Uber",
		Uri:     "https://eng.uber.com/mysql-migration/",
		Year:    2016,
		From: []string{
			"postgresql",
		},
		To: []string{
			"mysql",
		},
	},
	{
		Company: "Uber",
		Uri:     "https://www.infoq.com/articles/podcast-matt-ranney",
		Year:    2016,
		From: []string{
			"nodejs",
		},
		To: []string{
			"go",
			"java",
		},
	},
	{
		Company: "Uber",
		Uri:     "https://eng.uber.com/schemaless-rewrite/",
		Year:    2016,
		From: []string{
			"python",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "Digg",
		Uri:     "http://blog.digg.com/post/141552444676/making-the-switch-from-nodejs-to-golang",
		Year:    2016,
		From: []string{
			"nodejs",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "Google AdWords",
		Uri:     "https://news.dartlang.org/2016/03/the-new-adwords-ui-uses-dart-we-asked.html",
		Year:    2016,
		From: []string{
			"gwt",
		},
		To: []string{
			"angular",
			"dart",
		},
	},
	{
		Company: "Slack",
		Uri:     "https://jaxenter.com/php-hack-slack-168741.html",
		Year:    2016,
		From: []string{
			"php",
		},
		To: []string{
			"hack",
		},
	},
	{
		Company: "Sky",
		Uri:     "https://blog.couchbase.com/moving-from-oracle-to-couchbase/",
		Year:    2016,
		From: []string{
			"oracle",
		},
		To: []string{
			"couchbase",
		},
	},
	{
		Company: "Amadeus",
		Uri:     "https://www.computerweekly.com/news/450404428/Amadeus-turns-to-NoSQL-to-answer-complex-travel-questions",
		Year:    2016,
		From: []string{
			"oracle",
		},
		To: []string{
			"couchbase",
		},
	},
	{
		Company: "AppsFlyer",
		Uri:     "https://www.singlestore.com/blog/appsflyer-shares-journey-to-real-time-analytics-using-memsql/",
		Year:    2016,
		From: []string{
			"druid",
			"mongodb",
			"redis",
			"cassandra",
			"redshift",
		},
		To: []string{
			"singlestore",
		},
	},
	{
		Company: "Dropbox",
		Uri:     "https://www.datacenterknowledge.com/archives/2016/03/16/moving-away-from-aws-cloud-dropbox-isnt-an-anomaly-and-heres-why",
		Year:    2016,
		From: []string{
			"aws",
		},
		To: []string{
			"self-managed",
		},
	},
	{
		Company: "BleacherReport",
		Uri:     "https://www.techworld.com/apps-wearables/how-elixir-helped-bleacher-report-handle-8x-more-traffic-3653957/",
		Year:    2017,
		From: []string{
			"ruby",
		},
		To: []string{
			"elixir",
		},
	},
	{
		Company: "UpGuard",
		Uri:     "https://www.upguard.com/blog/our-experience-with-golang",
		Year:    2017,
		From: []string{
			"jruby",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "Movio",
		Uri:     "https://movio.co/blog/migrate-Scala-to-Go/",
		Year:    2017,
		From: []string{
			"scala",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "Discord",
		Uri:     "https://blog.discordapp.com/how-discord-stores-billions-of-messages-7fa6ec7ee4c7#.bv31n293t",
		Year:    2017,
		From: []string{
			"mongodb",
		},
		To: []string{
			"cassandra",
		},
	},
	{
		Company: "Uber",
		Uri:     "https://www.blogger.com/7https://eng.uber.com/distributed-tracing/",
		Year:    2017,
		From: []string{
			"riak",
			"solr",
		},
		To: []string{
			"cassandra",
		},
	},
	{
		Company: "Swat.io",
		Uri:     "https://web.archive.org/web/20181122212934/https://garage.socialisten.at/2017/04/how-swat-io-migrated-from-mysql-to-postgresql-in-2-years/",
		Year:    2017,
		From: []string{
			"mysql",
		},
		To: []string{
			"postgresql",
		},
	},
	{
		Company: "Poki",
		Uri:     "https://blog.poki.com/going-for-go-and-sticking-with-sql-a30faa42d643",
		Year:    2017,
		From: []string{
			"php",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "Stream",
		Uri:     "https://getstream.io/blog/switched-python-go/",
		Year:    2017,
		From: []string{
			"python",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "Baidu",
		Uri:     "https://www.cockroachlabs.com/case-studies/baidu/",
		Year:    2017,
		From: []string{
			"mysql",
		},
		To: []string{
			"cockroachdb",
		},
	},
	{
		Company: "Boxzilla",
		Uri:     "https://dannyvankooten.com/laravel-to-golang/",
		Year:    2017,
		From: []string{
			"php",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "GAEA",
		Uri:     "https://pingcap.medium.com/migration-from-mysql-to-tidb-to-handle-tens-of-millions-of-rows-of-data-per-day-f5b6e0e27d48",
		Year:    2017,
		From: []string{
			"mongodb",
			"mysql",
		},
		To: []string{
			"tidb",
		},
	},
	{
		Company: "Shippable",
		Uri:     "https://medium.com/thecobbles/why-we-moved-from-golang-to-nodejs-cecf66a47740",
		Year:    2017,
		From: []string{
			"go",
		},
		To: []string{
			"nodejs",
		},
	},
	{
		Company: "Shippable",
		Uri:     "https://dzone.com/articles/why-we-moved-from-nosql-mongodb-to-postgresql",
		Year:    2017,
		From: []string{
			"mongodb",
		},
		To: []string{
			"postgresql",
		},
	},
	{
		Company: "GeekyAnts",
		Uri:     "https://hackernoon.com/what-are-we-doing-with-googles-flutter-74ff29dd256a",
		Year:    2017,
		From: []string{
			"reactnative",
		},
		To: []string{
			"flutter",
		},
	},
	{
		Company: "Yuanfudao",
		Uri:     "https://pingcap.com/case-studies/tidb-in-yuanfudao",
		Year:    2017,
		From: []string{
			"mysql",
		},
		To: []string{
			"tidb",
		},
	},
	{
		Company: "Dropbox",
		Uri:     "https://dropbox.tech/frontend/the-great-coffeescript-to-typescript-migration-of-2017",
		Year:    2017,
		From: []string{
			"coffeescript",
		},
		To: []string{
			"typescript",
		},
	},
	{
		Company: "Manage",
		Uri:     "https://www.singlestore.com/blog/how-manage-accelerated-data-freshness-by-10x/",
		Year:    2017,
		From: []string{
			"mysql",
			"hadoop",
			"hive",
			"kafka",
		},
		To: []string{
			"singlestore",
		},
	},
	{
		Company: "BlueShift",
		Uri:     "https://www.singlestore.com/blog/blueshift-succeeding-with-real-time-analytics/",
		Year:    2017,
		From: []string{
			"mssqlserver",
		},
		To: []string{
			"singlestore",
		},
	},
	{
		Company: "Myntra",
		Uri:     "https://www.singlestore.com/blog/guest-post-real-time-big-data-ingestion-with-meterial/",
		Year:    2017,
		From: []string{
			"redshift",
		},
		To: []string{
			"singlestore",
		},
	},
	{
		Company: "TreeScale",
		Uri:     "https://medium.com/hackernoon/5-reasons-why-we-switched-from-python-to-go-4414d5f42690",
		Year:    2017,
		From: []string{
			"python",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "AppTree",
		Uri:     "https://www.youtube.com/watch?v=GpLb2XvKv20&list=PLOU2XLYxmsIIJr3vjxggY7yGcGO7i9BK5&index=4",
		Year:    2018,
		From: []string{
			"java",
			"kotlin",
			"objc",
			"swift",
		},
		To: []string{
			"dart",
		},
	},
	{
		Company: "SendGrid",
		Uri:     "https://stackshare.io/sendgrid/how-sendgrid-scaled-to-40-billion-emails-per-month",
		Year:    2018,
		From: []string{
			"perl",
			"python",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "Bukalapak",
		Uri:     "https://medium.com/inside-bukalapak/how-we-made-a-20-times-better-performance-microservice-part-1-e304df5b347b",
		Year:    2018,
		From: []string{
			"ruby",
			"mysql",
		},
		To: []string{
			"go",
			"mongodb",
			"elasticsearch",
			"cloudsql",
		},
	},
	{
		Company: "Centrifugo",
		Uri:     "https://medium.com/@fzambia/centrifugo-v2-0-released-built-on-top-of-new-real-time-messaging-library-for-go-language-b6ac034a6937",
		Year:    2018,
		From: []string{
			"python",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "TheGuardian",
		Uri:     "https://www.theguardian.com/info/2018/nov/30/bye-bye-mongo-hello-postgres",
		Year:    2018,
		From: []string{
			"mongodb",
		},
		To: []string{
			"postgresql",
		},
	},
	{
		Company: "CloudFlare",
		Uri:     "https://blog.cloudflare.com/http-analytics-for-6m-requests-per-second-using-clickhouse/",
		Year:    2018,
		From: []string{
			"citusdb",
		},
		To: []string{
			"clickhouse",
		},
	},
	{
		Company: "Github",
		Uri:     "https://github.blog/2018-09-06-removing-jquery-from-github-frontend/",
		Year:    2018,
		From: []string{
			"jquery",
		},
		To: []string{
			"javascript",
		},
	},
	{
		Company: "Meituan",
		Uri:     "https://pingcap.com/case-studies/migrating-from-mysql-to-a-scale-out-database-to-serve-our-290-million-monthly-users",
		Year:    2018,
		From: []string{
			"mysql",
		},
		To: []string{
			"tidb",
		},
	},
	{
		Company: "iQiyi",
		Uri:     "https://pingcap.com/case-studies/tidb-in-iqiyi",
		Year:    2018,
		From: []string{
			"mysql",
		},
		To: []string{
			"tidb",
		},
	},
	{
		Company: "Yiguo",
		Uri:     "https://pingcap.com/case-studies/hybrid-database-capturing-perishable-insights-at-yiguo",
		Year:    2018,
		From: []string{
			"hadoop",
			"mssqlserver",
		},
		To: []string{
			"tidb",
		},
	},
	{
		Company: "Mobike",
		Uri:     "https://pingcap.com/case-studies/tidb-in-mobike",
		Year:    2018,
		From: []string{
			"mysql",
		},
		To: []string{
			"tidb",
		},
	},
	{
		Company: "CashApp",
		Uri:     "https://developer.squareup.com/blog/sharding-cash/",
		Year:    2018,
		From: []string{
			"mysql",
		},
		To: []string{
			"vitess",
		},
	},
	{
		Company: "Bing",
		Uri:     "https://visualstudiomagazine.com/articles/2018/08/22/bing-net-core.aspx?m=1",
		Year:    2018,
		From: []string{
			"dotnet",
		},
		To: []string{
			"dotnetcore",
		},
	},
	{
		Company: "PixelJets",
		Uri:     "https://pixeljets.com/blog/clickhouse-as-a-replacement-for-elk-big-query-and-timescaledb/",
		Year:    2018,
		From: []string{
			"elsticsearch",
			"logstash",
			"kibana",
			"bigquery",
		},
		To: []string{
			"clickhouse",
		},
	},
	{
		Company: "Insite360",
		Uri:     "https://www.singlestore.com/blog/case-study-insite360-memsql-iot-cloud/",
		Year:    2018,
		From: []string{
			"redshift",
		},
		To: []string{
			"singlestore",
		},
	},
	{
		Company: "Areeba",
		Uri:     "https://www.singlestore.com/blog/areeba-case-study/",
		Year:    2018,
		From: []string{
			"hadoop",
			"mariadb",
		},
		To: []string{
			"singlestore",
		},
	},
	{
		Company: "Wag",
		Uri:     "https://www.singlestore.com/blog/wag-labs-case-study/",
		Year:    2018,
		From: []string{
			"mysql",
		},
		To: []string{
			"singlestore",
		},
	},
	{
		Company: "Fanatics",
		Uri:     "https://www.singlestore.com/blog/how-fanatics-powered-their-way-to-a-better-future/",
		Year:    2018,
		From: []string{
			"elasticsearch",
		},
		To: []string{
			"singlestore",
		},
	},
	{
		Company: "Pandora",
		Uri:     "https://www.singlestore.com/blog/pandora/",
		Year:    2018,
		From: []string{
			"hadoop",
		},
		To: []string{
			"singlestore",
		},
	},
	{
		Company: "Twitch",
		Uri:     "https://blog.twitch.tv/en/2022/04/12/breaking-the-monolith-at-twitch-part-2/",
		Year:    2018,
		From: []string{
			"ruby",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "Kairos",
		Uri:     "https://www.kairos.com/blog/php-to-go-how-we-boosted-api-performance-by-8x",
		Year:    2018,
		From: []string{
			"php",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "AppsFlyer",
		Uri:     "https://www.infoq.com/articles/api-gateway-clojure-golang",
		Year:    2019,
		From: []string{
			"clojure",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "Deliveroo",
		Uri:     "https://deliveroo.engineering/2019/02/14/moving-from-ruby-to-rust.html",
		Year:    2019,
		From: []string{
			"ruby",
		},
		To: []string{
			"rust",
		},
	},
	{
		Company: "Amazon",
		Uri:     "https://www.theregister.co.uk/2019/04/02/amazon_fulfilment_oracle_database",
		Year:    2019,
		From: []string{
			"oracle",
		},
		To: []string{
			"aurora_postgresql",
			"dynamodb",
		},
	},
	{
		Company: "2FintechGiants",
		Uri:     "https://www.youtube.com/watch?v=IG1E7O1rl-s",
		Year:    2019,
		From: []string{
			"oracle",
		},
		To: []string{
			"cockroachdb",
		},
	},
	{
		Company: "Codism",
		Uri:     "https://codism.io/why-we-migrated-from-python-to-golang",
		Year:    2019,
		From: []string{
			"python",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "Comcast",
		Uri:     "https://www.scylladb.com/tech-talk/sprinting-from-cassandra-to-scylladb/",
		Year:    2019,
		From: []string{
			"cassandra",
		},
		To: []string{
			"scylladb",
		},
	},
	{
		Company: "Zhihu",
		Uri:     "https://pingcap.com/case-studies/lesson-learned-from-queries-over-1.3-trillion-rows-of-data-within-milliseconds-of-response-time-at-zhihu",
		Year:    2019,
		From: []string{
			"mysql",
		},
		To: []string{
			"tidb",
		},
	},
	{
		Company: "JDCloud",
		Uri:     "https://pingcap.com/case-studies/lesson-learned-from-40-k-qps-and-20-billion-rows-of-data-in-a-single-scale-out-cluster",
		Year:    2019,
		From: []string{
			"mysql",
		},
		To: []string{
			"tidb",
		},
	},
	{
		Company: "NetEaseGames",
		Uri:     "https://pingcap.com/case-studies/why-we-chose-tidb-over-other-mysql-based-and-newsql-storage-solutions",
		Year:    2019,
		From: []string{
			"mysql",
		},
		To: []string{
			"tidb",
		},
	},
	{
		Company: "Xiaomi",
		Uri:     "https://pingcap.com/case-studies/tidb-in-xiaomi",
		Year:    2019,
		From: []string{
			"mysql",
		},
		To: []string{
			"tidb",
		},
	},
	{
		Company: "BookMyShow",
		Uri:     "https://pingcap.com/case-studies/tidb-in-bookmyshow",
		Year:    2019,
		From: []string{
			"mssqlserver",
		},
		To: []string{
			"tidb",
		},
	},
	{
		Company: "Trivago",
		Uri:     "https://searchitoperations.techtarget.com/news/252480628/HashiCorp-Nomad-vs-Kubernetes-matchup-intensifies-with-011",
		Year:    2019,
		From: []string{
			"kubernetes",
		},
		To: []string{
			"nomad",
		},
	},
	{
		Company: "Narvar",
		Uri:     "https://www.yugabyte.com/success-stories/narvar/",
		Year:    2019,
		From: []string{
			"dynamodb",
			"postgresql",
		},
		To: []string{
			"yugabytedb",
		},
	},
	{
		Company: "PlumeDesignsInc",
		Uri:     "https://www.yugabyte.com/success-stories/plume/",
		Year:    2019,
		From: []string{
			"mongodb",
		},
		To: []string{
			"yugabytedb",
		},
	},
	{
		Company: "SkyElectric",
		Uri:     "https://scylladb.medium.com/how-skyelectric-uses-scylla-to-power-its-smart-energy-platform-751f5b976d19",
		Year:    2019,
		From: []string{
			"mysql",
			"mongodb",
			"nodejs",
		},
		To: []string{
			"postgresql",
			"elasticsearch",
			"scylladb",
			"elixir",
		},
	},
	{
		Company: "Paytm",
		Uri:     "https://grafana.com/blog/2019/11/19/how-loki-helped-paytm-insider-save-75-of-logging-and-monitoring-costs/",
		Year:    2019,
		From: []string{
			"elasticsearch",
		},
		To: []string{
			"loki",
		},
	},
	{
		Company: "ContentSquare",
		Uri:     "https://www.youtube.com/watch?v=lwYSYMwpJOU",
		Year:    2019,
		From: []string{
			"elsticsearch",
			"logstash",
			"kibana",
		},
		To: []string{
			"clickhouse",
		},
	},
	{
		Company: "Sentry",
		Uri:     "https://blog.sentry.io/2019/05/16/introducing-snuba-sentrys-new-search-infrastructure",
		Year:    2019,
		From: []string{
			"postgresql",
			"redis",
		},
		To: []string{
			"clickhouse",
		},
	},
	{
		Company: "HashNode",
		Uri:     "https://engineering.hashnode.com/after-4-years-with-nginx-we-switched-to-caddy-here-is-why-cjxbv8eb2001ke8s1yl7ndroz",
		Year:    2019,
		From: []string{
			"nginx",
		},
		To: []string{
			"caddy",
		},
	},
	{
		Company: "CitizensBank",
		Uri:     "https://www.youtube.com/watch?v=YGujnkAV3pc",
		Year:    2019,
		From: []string{
			"kubernetes",
		},
		To: []string{
			"dockerswarm",
		},
	},
	{
		Company: "Medaxion",
		Uri:     "https://www.singlestore.com/blog/case-study-medaxion-analytics-medtech/",
		Year:    2019,
		From: []string{
			"mysql",
		},
		To: []string{
			"singlestore",
		},
	},
	{
		Company: "GoGuardian",
		Uri:     "https://www.singlestore.com/blog/case-study-goguardian-fast-analytics/",
		Year:    2019,
		From: []string{
			"druid",
		},
		To: []string{
			"singlestore",
		},
	},
	{
		Company: "AFortune50Company",
		Uri:     "https://www.singlestore.com/blog/case-study-hadoop-memsql-fortune-50/",
		Year:    2019,
		From: []string{
			"hadoop",
		},
		To: []string{
			"singlestore",
		},
	},
	{
		Company: "diwo",
		Uri:     "https://www.singlestore.com/blog/case-study-memsql-powering-ai-breakthroughs-at-diwo/",
		Year:    2019,
		From: []string{
			"redis",
			"cassandra",
		},
		To: []string{
			"singlestore",
		},
	},
	{
		Company: "KurtoSys",
		Uri:     "https://www.singlestore.com/blog/case-study-kurtosys-why-would-i-store-my-data-in-more-than-one-database/",
		Year:    2019,
		From: []string{
			"couchbase",
		},
		To: []string{
			"singlestore",
		},
	},
	{
		Company: "monday.com",
		Uri:     "https://www.singlestore.com/blog/case-study-mondaydotcom-bi/",
		Year:    2019,
		From: []string{
			"elasticsearch",
		},
		To: []string{
			"singlestore",
		},
	},
	{
		Company: "GumGum",
		Uri:     "https://medium.com/gumgum-tech/moving-to-amazon-dynamodb-from-hosted-cassandra-a-leap-towards-60-cost-saving-per-year-7eac8ac3bd55",
		Year:    2019,
		From: []string{
			"cassandra",
		},
		To: []string{
			"scylladb",
		},
	},
	{
		Company: "Insider",
		Uri:     "https://dev.to/ebaykann/how-we-moved-from-ruby-to-go-and-decrease-our-cost-by-1400-and-increased-response-time-by-500-2onj",
		Year:    2019,
		From: []string{
			"ruby",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "Codism",
		Uri:     "https://codism.io/why-we-migrated-from-python-to-golang/",
		Year:    2019,
		From: []string{
			"python",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "Twitch",
		Uri:     "https://blog.twitch.tv/en/2022/03/30/breaking-the-monolith-at-twitch/",
		Year:    2019,
		From: []string{
			"emberjs",
		},
		To: []string{
			"react",
		},
	},
	{
		Company: "Discord",
		Uri:     "https://blog.discordapp.com/why-discord-is-switching-from-go-to-rust",
		Year:    2020,
		From: []string{
			"go",
		},
		To: []string{
			"rust",
		},
	},
	{
		Company: "HSBC",
		Uri:     "https://diginomica.com/hsbc-moves-65-relational-databases-one-global-mongodb-database",
		Year:    2020,
		From: []string{
			"db2",
		},
		To: []string{
			"mongodb",
		},
	},
	{
		Company: "UnnamedUSWirelessCarrier",
		Uri:     "https://www.enterprisedb.com/resources/case-studies/us-wireless-carrier-migrates-100tb-oracle-database-edb-postgres-first-open",
		Year:    2020,
		From: []string{
			"oracle",
		},
		To: []string{
			"postgresql",
		},
	},
	{
		Company: "ListenBrainz",
		Uri:     "https://blog.metabrainz.org/2020/07/22/listenbrainz-moves-to-timescaledb/",
		Year:    2020,
		From: []string{
			"influxdb",
		},
		To: []string{
			"timescaledb",
		},
	},
	{
		Company: "Dropbox",
		Uri:     "https://dropbox.tech/infrastructure/how-we-migrated-dropbox-from-nginx-to-envoy",
		Year:    2020,
		From: []string{
			"nginx",
		},
		To: []string{
			"envoy",
		},
	},
	{
		Company: "Dropbox",
		Uri:     "https://dropbox.tech/infrastructure/rewriting-the-heart-of-our-sync-engine",
		Year:    2020,
		From: []string{
			"python",
		},
		To: []string{
			"rust",
		},
	},
	{
		Company: "Repustate",
		Uri:     "https://www.repustate.com/blog/migrating-entire-api-go-python/",
		Year:    2020,
		From: []string{
			"python",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "SimilarWeb",
		Uri:     "https://similarweb.engineering/moving-from-nodejs-to-go-doing-more-faster-for-less/",
		Year:    2020,
		From: []string{
			"nodejs",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "PayPal",
		Uri:     "https://go.dev/solutions/paypal/",
		Year:    2020,
		From: []string{
			"cpp",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "Discord",
		Uri:     "https://www.scylladb.com/press-release/discord-chooses-scylla-core-storage-layer/",
		Year:    2020,
		From: []string{
			"cassandra",
		},
		To: []string{
			"scylladb",
		},
	},
	{
		Company: "Weblove",
		Uri:     "https://betterprogramming.pub/why-we-moved-from-react-to-svelte-f20afb1dc5d5",
		Year:    2020,
		From: []string{
			"react",
		},
		To: []string{
			"svelte",
		},
	},
	{
		Company: "PalFish",
		Uri:     "https://pingcap.com/case-studies/embracing-newsql-why-we-chose-tidb-over-mongodb-and-mysql",
		Year:    2020,
		From: []string{
			"mongodb",
		},
		To: []string{
			"tidb",
		},
	},
	{
		Company: "Shopee",
		Uri:     "https://pingcap.com/case-studies/choosing-right-database-for-your-applications",
		Year:    2020,
		From: []string{
			"mysql",
		},
		To: []string{
			"tidb",
		},
	},
	{
		Company: "WeBank",
		Uri:     "https://pingcap.com/case-studies/how-we-reduced-batch-processing-time-by-58-percent-with-a-scale-out-mysql-alternative",
		Year:    2020,
		From: []string{
			"mysql",
		},
		To: []string{
			"tidb",
		},
	},
	{
		Company: "ChinaTelecomBestpay",
		Uri:     "https://pingcap.com/case-studies/how-we-process-data-five-times-more-efficiently-using-a-scale-out-mysql-alternative",
		Year:    2020,
		From: []string{
			"oracle",
		},
		To: []string{
			"tidb",
		},
	},
	{
		Company: "U-Next",
		Uri:     "https://pingcap.com/case-studies/running-a-scale-out-database-on-arm-as-mysql-alternative",
		Year:    2020,
		From: []string{
			"mysql",
		},
		To: []string{
			"tidb",
		},
	},
	{
		Company: "ZTOExpress",
		Uri:     "https://pingcap.com/case-studies/3x-it-efficiency-boost-use-a-scale-out-htap-database-for-near-real-time-analytics",
		Year:    2020,
		From: []string{
			"oracleexadata",
		},
		To: []string{
			"tidb",
		},
	},
	{
		Company: "Xiaohongshu",
		Uri:     "https://pingcap.com/case-studies/how-we-use-a-scale-out-htap-database-for-real-time-analytics-and-complex-queries",
		Year:    2020,
		From: []string{
			"mysql",
		},
		To: []string{
			"tidb",
		},
	},
	{
		Company: "BankOfBeijing",
		Uri:     "https://pingcap.com/case-studies/how-we-use-a-distributed-database-to-achieve-horizontal-scaling-without-downtime",
		Year:    2020,
		From: []string{
			"unknown",
		},
		To: []string{
			"tidb",
		},
	},
	{
		Company: "BankOfChina",
		Uri:     "https://pingcap.com/case-studies/how-bank-of-china-uses-a-scale-out-database-to-support-zabbix-monitoring-at-scale",
		Year:    2020,
		From: []string{
			"mysql",
		},
		To: []string{
			"tidb",
		},
	},
	{
		Company: "ZhuanZhuan",
		Uri:     "https://pingcap.com/case-studies/scale-out-database-powers-china-letgo-with-reduced-maintenance-costs",
		Year:    2020,
		From: []string{
			"mysql",
		},
		To: []string{
			"tidb",
		},
	},
	{
		Company: "BIGO",
		Uri:     "https://pingcap.com/case-studies/why-we-chose-an-htap-database-over-mysql-for-horizontal-scaling-and-complex-queries",
		Year:    2020,
		From: []string{
			"mysql",
		},
		To: []string{
			"tidb",
		},
	},
	{
		Company: "VNG/ZaloPay",
		Uri:     "https://pingcap.com/case-studies/zalopay-using-a-scale-out-mysql-alternative-to-serve-millions-of-users",
		Year:    2020,
		From: []string{
			"mysql",
		},
		To: []string{
			"tidb",
		},
	},
	{
		Company: "VIPKid",
		Uri:     "https://pingcap.com/case-studies/why-we-chose-a-distributed-sql-database-to-complement-mysql",
		Year:    2020,
		From: []string{
			"mysql",
		},
		To: []string{
			"tidb",
		},
	},
	{
		Company: "Kroger",
		Uri:     "https://vimeo.com/showcase/7581911/video/459492536",
		Year:    2020,
		From: []string{
			"postgresql",
		},
		To: []string{
			"yugabytedb",
		},
	},
	{
		Company: "Censys",
		Uri:     "https://vimeo.com/showcase/7581911/video/459492564",
		Year:    2020,
		From: []string{
			"postgresql",
		},
		To: []string{
			"yugabytedb",
		},
	},
	{
		Company: "Justuno",
		Uri:     "https://vimeo.com/showcase/7581911/video/459492578",
		Year:    2020,
		From: []string{
			"mssqlserver",
			"cassandra",
		},
		To: []string{
			"yugabytedb",
		},
	},
	{
		Company: "Xignite",
		Uri:     "https://www.yugabyte.com/success-stories/xignite/",
		Year:    2020,
		From: []string{
			"mssqlserver",
		},
		To: []string{
			"yugabytedb",
		},
	},
	{
		Company: "Voiceland",
		Uri:     "https://www.yugabyte.com/success-stories/voiceland/",
		Year:    2020,
		From: []string{
			"postgresql",
			"mssqlserver",
		},
		To: []string{
			"yugabytedb",
		},
	},
	{
		Company: "Jutsuno",
		Uri:     "https://www.yugabyte.com/success-stories/justuno/",
		Year:    2020,
		From: []string{
			"cassandra",
			"neo4j",
			"mssqlserver",
			"cockroachdb",
		},
		To: []string{
			"yugabytedb",
		},
	},
	{
		Company: "Manetu",
		Uri:     "https://www.yugabyte.com/success-stories/manetu/",
		Year:    2020,
		From: []string{
			"cassandra",
		},
		To: []string{
			"yugabytedb",
		},
	},
	{
		Company: "BRIKL",
		Uri:     "https://medium.com/yugabyte/distributed-sql-summit-recap-a-migration-journey-from-amazon-dynamodb-to-yugabytedb-and-hasura-507189cd9074",
		Year:    2020,
		From: []string{
			"dynamodb",
		},
		To: []string{
			"yugabytedb",
			"hasura",
		},
	},
	{
		Company: "StadiaMaps",
		Uri:     "https://scylladb.medium.com/stadia-maps-using-scylla-to-serve-maps-in-milliseconds-640c0e4923b",
		Year:    2020,
		From: []string{
			"cockroachdb",
		},
		To: []string{
			"scylladb",
		},
	},
	{
		Company: "Roblox",
		Uri:     "https://www.youtube.com/watch?v=6xOUbJ7wnP8",
		Year:    2020,
		From: []string{
			"windows",
		},
		To: []string{
			"linux",
			"nomad",
		},
	},
	{
		Company: "Rekki",
		Uri:     "https://dev.to/rekki/why-we-killed-elixir-3np",
		Year:    2020,
		From: []string{
			"elixir",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "Tubber",
		Uri:     "https://roelofjanelsinga.com/articles/the-impact-of-migrating-from-php-to-golang/",
		Year:    2020,
		From: []string{
			"php",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "EBay",
		Uri:     "https://tech.ebayinc.com/engineering/ou-online-analytical-processing/",
		Year:    2020,
		From: []string{
			"apachedruid",
		},
		To: []string{
			"clickhouse",
		},
	},
	{
		Company: "Shippo",
		Uri:     "https://goshippo.com/blog/the-road-to-aurora-postgresql-with-near-zero-downtime/",
		Year:    2020,
		From: []string{
			"postgresql",
		},
		To: []string{
			"aurora",
		},
	},
	{
		Company: "Nucleus",
		Uri:     "https://www.singlestore.com/customers/nucleus/",
		Year:    2020,
		From: []string{
			"mariadb",
		},
		To: []string{
			"singlestore",
		},
	},
	{
		Company: "AMajorOilAndGasCompany",
		Uri:     "https://www.singlestore.com/blog/memsql-improves-financial-operations-for-a-major-oil-and-gas-company/",
		Year:    2020,
		From: []string{
			"elasticsearch",
		},
		To: []string{
			"singlestore",
		},
	},
	{
		Company: "TapJoy",
		Uri:     "https://www.singlestore.com/blog/tapjoy-moving-to-memsql/",
		Year:    2020,
		From: []string{
			"mysql",
		},
		To: []string{
			"singlestore",
		},
	},
	{
		Company: "ARNES",
		Uri:     "https://github.com/VictoriaMetrics/VictoriaMetrics/commit/8064775c02d5b24fcff65eb3336f6b67027ecb24#diff-ff50e6019204ee4ddf2a9d295a28f1ab679d9bb56d17cc30057ce42566034a9f",
		Year:    2020,
		From: []string{
			"influxdb",
		},
		To: []string{
			"victoriametrics",
		},
	},
	{
		Company: "DevOpsProdigy",
		Uri:     "https://devopsprodigy.com/blog/chose-the-right-time-series-database/",
		Year:    2020,
		From: []string{
			"mysql",
		},
		To: []string{
			"clickhouse",
		},
	},
	{
		Company: "MyWorld",
		Uri:     "https://www.cockroachlabs.com/blog/cassandra-to-cockroachdb/",
		Year:    2020,
		From: []string{
			"cassandra",
		},
		To: []string{
			"cockroachdb",
		},
	},
	{
		Company: "PlotProjects",
		Uri:     "https://www.plotprojects.com/blog/why-we-use-postgresql-and-slick/",
		Year:    2020,
		From: []string{
			"mongodb",
		},
		To: []string{
			"postgresql",
		},
	},
	{
		Company: "SpaceCloud",
		Uri:     "https://blog.space-cloud.io/posts/why-we-moved-from-grpc-to-graphql/",
		Year:    2020,
		From: []string{
			"grpc",
		},
		To: []string{
			"graphql",
		},
	},
	{
		Company: "Sudo",
		Uri:     "https://betterprogramming.pub/why-we-moved-from-react-to-svelte-f20afb1dc5d5",
		Year:    2020,
		From: []string{
			"react",
		},
		To: []string{
			"svelte",
		},
	},
	{
		Company: "LoginRadius",
		Uri:     "https://www.loginradius.com/blog/engineering/a-journey-from-node-to-golang/",
		Year:    2020,
		From: []string{
			"nodejs",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "KhanAcademy",
		Uri:     "https://youtu.be/je9bC3DZ6tg",
		Year:    2021,
		From: []string{
			"python",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "ownCloud",
		Uri:     "https://www.heise.de//news/ownCloud-Infinite-Scale-Go-statt-PHP-Microservices-statt-LAMP-5029244.html",
		Year:    2021,
		From: []string{
			"php",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "Percona",
		Uri:     "https://www.percona.com/blog/2020/12/23/observations-on-better-resource-usage-with-percona-monitoring-and-management-v2-12-0/",
		Year:    2021,
		From: []string{
			"prometheus",
		},
		To: []string{
			"victoriametrics",
		},
	},
	{
		Company: "Aluma",
		Uri:     "https://aluma.io/resources/blog/switching-from-c-to-go-for-backend-development",
		Year:    2021,
		From: []string{
			"c#",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "Olery",
		Uri:     "https://developer.olery.com/blog/goodbye-mongodb-hello-postgresql/",
		Year:    2021,
		From: []string{
			"mongodb",
		},
		To: []string{
			"postgresql",
		},
	},
	{
		Company: "NitroKey",
		Uri:     "https://www.nitrokey.com/news/2021/nextbox-why-we-decided-and-against-ubuntu-core",
		Year:    2021,
		From: []string{
			"ubuntu_core",
		},
		To: []string{
			"debian",
		},
	},
	{
		Company: "Slack",
		Uri:     "https://slack.engineering/migrating-millions-of-concurrent-websockets-to-envoy/",
		Year:    2021,
		From: []string{
			"haproxy",
		},
		To: []string{
			"envoy",
		},
	},
	{
		Company: "Github",
		Uri:     "https://fhnjxazz4qgumnpam4q7tg4vxi-adwhj77lcyoafdy-github-blog.translate.goog/jp/2021-03-19-improving-large-monorepo-performance-on-github/",
		Year:    2021,
		From: []string{
			"ruby",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "GumGum",
		Uri:     "https://www.youtube.com/watch?v=RR5j_1HV7ng",
		Year:    2021,
		From: []string{
			"cassandra",
		},
		To: []string{
			"scylladb",
		},
	},
	{
		Company: "SvelteKit",
		Uri:     "https://svelte.dev/blog/sveltekit-beta",
		Year:    2021,
		From: []string{
			"rollup",
			"snowpack",
		},
		To: []string{
			"vite",
		},
	},
	{
		Company: "ReplIt",
		Uri:     "https://blog.replit.com/vite",
		Year:    2021,
		From: []string{
			"webpack",
		},
		To: []string{
			"vite",
		},
	},
	{
		Company: "GitPod",
		Uri:     "https://www.gitpod.io/blog/from-gatsby-to-svelte",
		Year:    2021,
		From: []string{
			"react",
			"gatsby",
		},
		To: []string{
			"sveltekit",
		},
	},
	{
		Company: "NinjaVan",
		Uri:     "https://pingcap.com/case-studies/choose-a-mysql-alternative-over-vitess-and-crdb-to-scale-out-our-databases-on-k8s",
		Year:    2021,
		From: []string{
			"mysql",
		},
		To: []string{
			"tidb",
		},
	},
	{
		Company: "Meituan",
		Uri:     "https://pingcap.com/case-studies/how-we-use-a-mysql-alternative-to-avoid-sharding-and-provide-strong-consistency",
		Year:    2021,
		From: []string{
			"mysql",
		},
		To: []string{
			"tidb",
		},
	},
	{
		Company: "AutoHome",
		Uri:     "https://pingcap.com/case-studies/reduce-real-time-query-latency-from-0.5s-to-0.01s-with-scale-out-htap-database",
		Year:    2021,
		From: []string{
			"mssqlserver",
		},
		To: []string{
			"tidb",
		},
	},
	{
		Company: "58.com",
		Uri:     "https://pingcap.com/case-studies/no-sharding-no-etl-use-scale-out-mysql-alternative-to-store-160-tb-of-data",
		Year:    2021,
		From: []string{
			"mysql",
		},
		To: []string{
			"tidb",
		},
	},
	{
		Company: "Chehaoduo",
		Uri:     "https://pingcap.com/case-studies/top-car-trading-platform-chooses-scale-out-database-as-mysql-alternative",
		Year:    2021,
		From: []string{
			"mysql",
		},
		To: []string{
			"tidb",
		},
	},
	{
		Company: "KunLun",
		Uri:     "https://pingcap.com/case-studies/empowering-your-gaming-application-with-a-scale-out-newsql-database",
		Year:    2021,
		From: []string{
			"mysql",
		},
		To: []string{
			"tidb",
		},
	},
	{
		Company: "GiG",
		Uri:     "https://www.youtube.com/watch?v=vkYvuIs1KcU&list=PLWhC0zeznqkkNYzcvHEfZ8hly3Cu9ojKk&index=8",
		Year:    2021,
		From: []string{
			"redshift",
		},
		To: []string{
			"clickhouse",
		},
	},
	{
		Company: "Brex",
		Uri:     "https://medium.com/brexeng/building-backend-services-with-kotlin-7c8410795e4b",
		Year:    2021,
		From: []string{
			"elixir",
		},
		To: []string{
			"kotlin",
		},
	},
	{
		Company: "Storj",
		Uri:     "https://thenewstack.io/why-the-storj-cloud-storage-service-switched-to-golang/",
		Year:    2021,
		From: []string{
			"python",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "CoinBase",
		Uri:     "https://blog.coinbase.com/announcing-coinbases-successful-transition-to-react-native-af4c591df971",
		Year:    2021,
		From: []string{
			"native",
		},
		To: []string{
			"reactnative",
		},
	},
	{
		Company: "Archive.org",
		Uri:     "https://www.youtube.com/watch?v=1n1gPMxg8bg",
		Year:    2021,
		From: []string{
			"kubernetes",
		},
		To: []string{
			"nomad",
			"consul",
		},
	},
	{
		Company: "Koyeb",
		Uri:     "https://www.koyeb.com/blog/the-koyeb-serverless-engine-from-kubernetes-to-nomad-firecracker-and-kuma",
		Year:    2021,
		From: []string{
			"kubernetes",
		},
		To: []string{
			"nomad",
			"firecracker",
			"kuma",
		},
	},
	{
		Company: "AccelByte",
		Uri:     "https://www.youtube.com/watch?v=-Zwr0CuPoCQ",
		Year:    2021,
		From: []string{
			"kubernetes",
		},
		To: []string{
			"nomad",
			"consul",
		},
	},
	{
		Company: "Expedia",
		Uri:     "https://scylladb.medium.com/expedia-group-our-migration-journey-to-scylla-12cdfa820442",
		Year:    2021,
		From: []string{
			"cassandra",
		},
		To: []string{
			"scylladb",
		},
	},
	{
		Company: "Grab",
		Uri:     "https://scylladb.medium.com/at-scylla-summit-2021-we-were-joined-by-two-members-of-the-engineering-team-at-grab-chao-wang-and-7a688d36a1f0",
		Year:    2021,
		From: []string{
			"cassandra",
		},
		To: []string{
			"scylladb",
		},
	},
	{
		Company: "ReadyForSky",
		Uri:     "https://medium.com/@tarantool/can-tarantool-beat-redis-in-iot-703bbc781d6e",
		Year:    2021,
		From: []string{
			"redis",
		},
		To: []string{
			"tarantool",
		},
	},
	{
		Company: "Conductor",
		Uri:     "https://thenewstack.io/conductor-why-we-migrated-from-kubernetes-to-nomad/",
		Year:    2021,
		From: []string{
			"kubernetes",
		},
		To: []string{
			"nomad",
		},
	},
	{
		Company: "Uber",
		Uri:     "https://eng.uber.com/logging/",
		Year:    2021,
		From: []string{
			"elsticsearch",
			"logstash",
			"kibana",
		},
		To: []string{
			"clickhouse",
		},
	},
	{
		Company: "ClickHouse",
		Uri:     "https://www.youtube.com/watch?v=abhcCRW09Ac",
		Year:    2021,
		From: []string{
			"zookeeper",
		},
		To: []string{
			"clickhousekeeper",
			"nuraft",
		},
	},
	{
		Company: "Wallaroo",
		Uri:     "https://www.wallaroo.ai/blog/wallaroo-move-to-rust",
		Year:    2021,
		From: []string{
			"pony",
		},
		To: []string{
			"rust",
		},
	},
	{
		Company: "ContentServ",
		Uri:     "https://www.youtube.com/watch?v=P-HBEDzNuWg",
		Year:    2021,
		From: []string{
			"cassandra",
		},
		To: []string{
			"scylladb",
		},
	},
	{
		Company: "Etsy",
		Uri:     "https://codeascraft.com/2021/11/08/etsys-journey-to-typescript/",
		Year:    2021,
		From: []string{
			"javascript",
		},
		To: []string{
			"typescript",
		},
	},
	{
		Company: "JD.com",
		Uri:     "https://pingcap.com/case-studies/8x-system-performance-boost-why-we-migrated-from-mysql-to-newsql-database",
		Year:    2021,
		From: []string{
			"mysql",
		},
		To: []string{
			"tidb",
		},
	},
	{
		Company: "HuyaLive",
		Uri:     "https://pingcap.com/case-studies/how-we-scale-out-databases-and-get-big-data-queries-6x-faster-with-a-mysql-alternative",
		Year:    2021,
		From: []string{
			"mysql",
		},
		To: []string{
			"tidb",
		},
	},
	{
		Company: "LevelUpTutorials",
		Uri:     "https://www.youtube.com/watch?v=ezk6qAIXe68",
		Year:    2021,
		From: []string{
			"react",
		},
		To: []string{
			"svelte",
		},
	},
	{
		Company: "Voucherify",
		Uri:     "https://www.voucherify.io/tech/how-we-moved-from-mongodb-to-postgres-without-downtime-and-cut-our-costs-by-30",
		Year:    2021,
		From: []string{
			"mongodb",
		},
		To: []string{
			"postgresql",
		},
	},
	{
		Company: "Fathom",
		Uri:     "https://www.singlestore.com/customers/fathom/",
		Year:    2021,
		From: []string{
			"mysql",
			"redis",
			"dynamodb",
		},
		To: []string{
			"singlestore",
		},
	},
	{
		Company: "GrooveX",
		Uri:     "https://github.com/VictoriaMetrics/VictoriaMetrics/commit/906fca9e88c74861c5486f4cdcd1b9caaba3e790#diff-ff50e6019204ee4ddf2a9d295a28f1ab679d9bb56d17cc30057ce42566034a9f",
		Year:    2021,
		From: []string{
			"thanos",
		},
		To: []string{
			"victoriametrics",
		},
	},
	{
		Company: "Percona",
		Uri:     "https://github.com/VictoriaMetrics/VictoriaMetrics/commit/ede9dd43e84034c1a8fc1080fabfe374b0f9155e#diff-ff50e6019204ee4ddf2a9d295a28f1ab679d9bb56d17cc30057ce42566034a9f",
		Year:    2021,
		From: []string{
			"prometheus",
		},
		To: []string{
			"victoriametrics",
		},
	},
	{
		Company: "SimilarWeb",
		Uri:     "https://medium.com/similarweb-engineering/moving-from-nodejs-to-go-doing-more-faster-for-less-1dfe45fd6b6b",
		Year:    2021,
		From: []string{
			"nodejs",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "DoorDash",
		Uri:     "https://doordash.engineering/2021/05/04/migrating-from-python-to-kotlin-for-our-backend-services",
		Year:    2021,
		From: []string{
			"python",
		},
		To: []string{
			"kotlin",
		},
	},
	{
		Company: "WeWatch",
		Uri:     "https://jerseyfonseca.com/blogs/mongodb-to-postgresql-migration",
		Year:    2021,
		From: []string{
			"mongodb",
		},
		To: []string{
			"postgresql",
		},
	},
	{
		Company: "CodeSignal",
		Uri:     "https://codesignal.com/blog/engineering/why-we-moved-from-mongodb-meteor-observers-to-redis/",
		Year:    2021,
		From: []string{
			"mongodb",
			"meteor",
		},
		To: []string{
			"redis",
		},
	},
	{
		Company: "Sequoia",
		Uri:     "https://rockset.com/blog/sequoia-capital-elasticsearch-to-rockset/",
		Year:    2021,
		From: []string{
			"elasticsearch",
		},
		To: []string{
			"rockset",
		},
	},
	{
		Company: "TimeFlow",
		Uri:     "https://news.knowledia.com/US/en/articles/why-we-moved-from-druid-to-clickhouse-74c743bf41a9a7fc2bc7c9c2ab37b4e1d5c966b7",
		Year:    2021,
		From: []string{
			"druid",
		},
		To: []string{
			"clickhouse",
		},
	},
	{
		Company: "ToolJet",
		Uri:     "https://blog.tooljet.com/how-we-migrated-tooljet-server-from-ruby-to-node-js/",
		Year:    2021,
		From: []string{
			"ruby",
		},
		To: []string{
			"nodejs",
		},
	},
	{
		Company: "Dream11",
		Uri:     "https://aerospike.com/resources/videos/dream11-architecting-scale-with-aerospike/",
		Year:    2021,
		From: []string{
			"elasticache_redis",
		},
		To: []string{
			"aerospike",
		},
	},
	{
		Company: "AirTel",
		Uri:     "https://aerospike.com/customers/airtel/",
		Year:    2021,
		From: []string{
			"oracle_timesten",
		},
		To: []string{
			"aerospike",
		},
	},
	{
		Company: "PostHog",
		Uri:     "https://www.youtube.com/watch?v=6IwLWEx_mg4",
		Year:    2021,
		From: []string{
			"postgresql",
		},
		To: []string{
			"clickhouse",
		},
	},
	{
		Company: "Factors.ai",
		Uri:     "https://www.singlestore.com/resources/webinar-turbocharge-your-open-source-db-to-drive-100x-faster-performance-2022-01/",
		Year:    2022,
		From: []string{
			"postgresql",
		},
		To: []string{
			"singlestore",
		},
	},
	{
		Company: "Prismatic.io",
		Uri:     "https://faun.pub/why-we-moved-from-lambda-to-ecs-b84674f31869",
		Year:    2022,
		From: []string{
			"aws_lambda",
		},
		To: []string{
			"amazon_ecs",
		},
	},
	{
		Company: "Fleet",
		Uri:     "https://blog.fleetdm.com/saving-over-100x-on-egress-switching-from-aws-to-hetzner-169888bd6650",
		Year:    2022,
		From: []string{
			"aws",
		},
		To: []string{
			"hetzner",
		},
	},
	{
		Company: "Ntop",
		Uri:     "https://www.ntop.org/ntop/historical-traffic-analysis-at-scale-using-clickhouse-with-ntopng/?utm_campaign=Social%20media&utm_content=194902279&utm_medium=social&utm_source=twitter&hss_channel=tw-3894792263",
		Year:    2022,
		From: []string{
			"nindex",
		},
		To: []string{
			"clickhouse",
		},
	},
	{
		Company: "Etsy",
		Uri:     "https://www.infoq.com/news/2022/01/etst-migration-from-react-preact/",
		Year:    2022,
		From: []string{
			"react",
		},
		To: []string{
			"preact",
		},
	},
	{
		Company: "ReviewBunny",
		Uri:     "https://reviewbunny.app/blog/dont-make-me-think-or-why-i-switched-to-rails-from-javascript-spas",
		Year:    2022,
		From: []string{
			"react",
		},
		To: []string{
			"rubyonrails",
		},
	},
	{
		Company: "AntMoney",
		Uri:     "https://www.singlestore.com/resources/webinar-migrating-from-postgresql-to-drive-20-100x-faster-performance-2022-02/",
		Year:    2022,
		From: []string{
			"postgresql",
		},
		To: []string{
			"singlestore",
		},
	},
	{
		Company: "CaptainMetrics",
		Uri:     "https://www.singlestore.com/resources/webinar-captain-metrics-why-we-ditched-mongodb-2022-03/",
		Year:    2022,
		From: []string{
			"mongodb",
		},
		To: []string{
			"singlestore",
		},
	},
	{
		Company: "SeventhSense",
		Uri:     "https://www.youtube.com/watch?v=8JFiDx6vLKQ",
		Year:    2022,
		From: []string{
			"amazon_kinesis",
			"kafka",
		},
		To: []string{
			"redpanda",
		},
	},
	{
		Company: "Hashura",
		Uri:     "https://nhost.io/blog/hasura-storage-in-go-5x-performance-increase-and-40-percent-less-ram",
		Year:    2022,
		From: []string{
			"nodejs",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "Kiwi",
		Uri:     "https://www.youtube.com/watch?v=YFm62QHIW1E",
		Year:    2022,
		From: []string{
			"postgresql",
			"cassandra",
		},
		To: []string{
			"scylladb",
		},
	},
	{
		Company: "Rakuten",
		Uri:     "https://www.youtube.com/watch?v=BS13KnN29AU",
		Year:    2022,
		From: []string{
			"cassandra",
		},
		To: []string{
			"scylladb",
		},
	},
	{
		Company: "Wipro",
		Uri:     "https://www.youtube.com/watch?v=rHJUf6qGCCM",
		Year:    2022,
		From: []string{
			"hbase",
		},
		To: []string{
			"aerospike",
		},
	},
	{
		Company: "nhost",
		Uri:     "https://nhost.io/blog/individual-postgres-instances",
		Year:    2022,
		From: []string{
			"postgresql",
		},
		To: []string{
			"postgresql",
			"kubernetes",
		},
	},
	{
		Company: "Prerender.io",
		Uri:     "https://levelup.gitconnected.com/how-we-reduced-our-annual-server-costs-by-80-from-1m-to-200k-by-moving-away-from-aws-2b98cbd21b46",
		Year:    2022,
		From: []string{
			"aws",
		},
		To: []string{
			"self-managed",
		},
	},
	{
		Company: "Umes",
		Uri:     "https://itnext.io/why-our-company-replaced-golang-graphql-with-typescript-prisma-trpc-ef56aaaa1c8c",
		Year:    2022,
		From: []string{
			"go",
			"graphql",
		},
		To: []string{
			"typescript",
			"prisma",
			"trpc",
		},
	},
	{
		Company: "K-Optional",
		Uri:     "https://koptional.com/article/why-we%E2%80%99re-moving-away-from-firebase",
		Year:    2022,
		From: []string{
			"firebase",
		},
		To: []string{
			"supabase",
		},
	},
	{
		Company: "Contexte",
		Uri:     "https://www.youtube.com/watch?v=3GObi93tjZI",
		Year:    2022,
		From: []string{
			"react",
		},
		To: []string{
			"htmx",
		},
	},
	{
		Company: "Rakuten",
		Uri:     "https://www.telecomtv.com/content/digital-platforms-services/rakuten-dumps-red-hat-turns-to-true-open-source-linux-os-45803/",
		Year:    2022,
		From: []string{
			"redhat",
		},
		To: []string{
			"rockylinux",
		},
	},
	{
		Company: "NucleusData",
		Uri:     "https://www.singlestore.com/resources/webinar-how-nucleus-security-drove-4X-throughput-with-distributed-sql-2022-12/",
		Year:    2022,
		From: []string{
			"mariadb",
		},
		To: []string{
			"singlestore",
		},
	},
	{
		Company: "Rolebase",
		Uri:     "https://github.com/nhost/nhost/discussions/1121",
		Year:    2022,
		From: []string{
			"firebase",
		},
		To: []string{
			"nhost",
		},
	},
	{
		Company: "Jamf",
		Uri:     "https://medium.com/jamf-engineering/making-our-lives-easier-by-rewriting-a-global-dns-gateway-in-go-3e39cdcdae83",
		Year:    2022,
		From: []string{
			"java",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "TheirStack.com",
		Uri:     "https://www.tinybird.co/blog-posts/text-search-at-scale-with-clickhouse",
		Year:    2022,
		From: []string{
			"postgresql",
		},
		To: []string{
			"clickhouse",
		},
	},
	{
		Company: "Matrix",
		Uri:     "https://www.synadia.com/blog/dendrite-project-moves-from-kafka-to-nats",
		Year:    2022,
		From: []string{
			"kafka",
		},
		To: []string{
			"nats",
		},
	},
	{
		Company: "CloudFlare",
		Uri:     "https://blog.cloudflare.com/log-analytics-using-clickhouse/",
		Year:    2022,
		From: []string{
			"elasticsearch",
		},
		To: []string{
			"clickhouse",
		},
	},
	{
		Company: "HiFi",
		Uri:     "https://clickhouse.com/blog/hifis-migration-from-bigquery-to-clickhouse",
		Year:    2023,
		From: []string{
			"bigquery",
		},
		To: []string{
			"clickhouse",
		},
	},
	{
		Company: "RisingWaveLabs",
		Uri:     "https://www.risingwave-labs.com/blog/building-a-cloud-database-from-scratch-why-we-moved-from-cpp-to-rust/",
		Year:    2023,
		From: []string{
			"cpp",
		},
		To: []string{
			"rust",
		},
	},
	{
		Company: "Vercel",
		Uri:     "https://vercel.com/blog/turborepo-migration-go-rust",
		Year:    2023,
		From: []string{
			"go",
		},
		To: []string{
			"rust",
		},
	},
	{
		Company: "Discord",
		Uri:     "https://discord.com/blog/how-discord-stores-trillions-of-messages",
		Year:    2023,
		From: []string{
			"cassandra",
		},
		To: []string{
			"scylladb",
		},
	},
	{
		Company: "RocLang",
		Uri:     "https://zackoverflow.dev/writing/unsafe-rust-vs-zig/",
		Year:    2023,
		From: []string{
			"rust",
		},
		To: []string{
			"zig",
		},
	},
	{
		Company: "37Signals",
		Uri:     "https://dev.37signals.com/bringing-our-apps-back-home/",
		Year:    2023,
		From: []string{
			"kubernetes",
		},
		To: []string{
			"mrsk",
		},
	},
	{
		Company: "CultureAmp",
		Uri:     "https://kevinyank.com/posts/on-endings-why-how-we-retired-elm-at-culture-amp/",
		Year:    2023,
		From: []string{
			"elm",
		},
		To: []string{
			"react",
		},
	},
	{
		Company: "ApiGear",
		Uri:     "https://www.reddit.com/r/golang/comments/12fvi7e/comment/jfjg4as/?utm_source=reddit&utm_medium=web2x&context=3",
		Year:    2023,
		From: []string{
			"nodejs",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "Prequel",
		Uri:     "https://www.prequel.co/blog/sql-maxis-why-we-ditched-rabbitmq-and-replaced-it-with-a-postgres-queue",
		Year:    2023,
		From: []string{
			"rabbitmq",
		},
		To: []string{
			"postgresql",
		},
	},
	{
		Company: "Trendyol",
		Uri:     "https://medium.com/trendyol-tech/new-winner-of-kafka-consumers-scala-to-go-journey-604c6bdd7041",
		Year:    2023,
		From: []string{
			"scala",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "AmazonPrimeVideo",
		Uri:     "https://www.primevideotech.com/video-streaming/scaling-up-the-prime-video-audio-video-monitoring-service-and-reducing-costs-by-90",
		Year:    2023,
		From: []string{
			"aws_lambda",
		},
		To: []string{
			"amazon_ec2",
			"amazon_ecs",
		},
	},
	{
		Company: "Svelte",
		Uri:     "https://thenewstack.io/rich-harris-talks-sveltekit-and-whats-next-for-svelte/",
		Year:    2023,
		From: []string{
			"typescript",
		},
		To: []string{
			"javascript",
			"jsdoc",
		},
	},
	{
		Company: "ValTown",
		Uri:     "https://blog.val.town/blog/migrating-from-supabase",
		Year:    2023,
		From: []string{
			"supabase",
		},
		To: []string{
			"postgresql",
			"render",
		},
	},
	{
		Company: "ChessCraft",
		Uri:     "https://blog.stuartspence.ca/2023-05-goodbye-mongo.html",
		Year:    2023,
		From: []string{
			"mongodb",
		},
		To: []string{
			"postgresql",
			"redis",
		},
	},
	{
		Company: "Fly.io",
		Uri:     "https://fly.io/blog/carving-the-scheduler-out-of-our-orchestrator/",
		Year:    2023,
		From: []string{
			"nomad",
		},
		To: []string{
			"flyd",
		},
	},
	{
		Company: "CodeDamn",
		Uri:     "https://codedamn.com/news/product/dont-use-prisma",
		Year:    2023,
		From: []string{
			"prisma",
			"mongodb",
		},
		To: []string{
			"postgresql",
		},
	},
	{
		Company: "LinkedIn",
		Uri:     "https://www.infoq.com/news/2023/07/linkedin-protocol-buffers-restli/",
		Year:    2023,
		From: []string{
			"json",
		},
		To: []string{
			"protobuf",
		},
	},
	{
		Company: "Bookshop.org",
		Uri:     "https://encore.dev/customers/bookshop",
		Year:    2023,
		From: []string{
			"rubyonrails",
		},
		To: []string{
			"go",
			"encore",
		},
	},
	{
		Company: "Lizza",
		Uri:     "https://encore.dev/customers/lizza",
		Year:    2023,
		From: []string{
			"firebase",
		},
		To: []string{
			"encore",
		},
	},
	{
		Company: "InfluxData",
		Uri:     "https://www.influxdata.com/blog/the-plan-for-influxdb-3-0-open-source/",
		Year:    2023,
		From: []string{
			"go",
		},
		To: []string{
			"rust",
		},
	},
	{
		Company: "FileWave",
		Uri:     "https://www.synadia.com/blog/filewave-why-we-moved-from-zeromq-to-nats",
		Year:    2023,
		From: []string{
			"zeromq",
		},
		To: []string{
			"nats",
		},
	},
	{
		Company: "StackOverflow",
		Uri:     "https://stackoverflow.blog/2023/10/31/why-stack-overflow-is-embracing-svelte/",
		Year:    2023,
		From: []string{
			"jquery",
		},
		To: []string{
			"svelte",
		},
	},
	{
		Company: "NicelandVPN",
		Uri:     "https://discord.com/channels/1042734330029547630/1149257802611699743/1149257802611699743",
		Year:    2023,
		From: []string{
			"electron",
		},
		To: []string{
			"wails.io",
		},
	},
	{
		Company: "DoorDash",
		Uri:     "https://thenewstack.io/how-doordash-migrated-from-aurora-postgres-to-cockroachdb/",
		Year:    2023,
		From: []string{
			"aurora_postgresql",
		},
		To: []string{
			"cockroachdb",
		},
	},
	{
		Company: "Levenue",
		Uri:     "https://www.youtube.com/watch?v=zB9tEQYLPL4",
		Year:    2023,
		From: []string{
			"sveltekit",
		},
		To: []string{
			"go",
			"htmx",
		},
	},
	{
		Company: "Lyft",
		Uri:     "https://eng.lyft.com/druid-deprecation-and-clickhouse-adoption-at-lyft-120af37651fd?gi=187126e686a7",
		Year:    2023,
		From: []string{
			"druid",
		},
		To: []string{
			"kafka",
			"clickhouse",
		},
	},
	{
		Company: "Tencent",
		Uri:     "https://www.reddit.com/r/datascience/comments/11mgmf1/tencent_data_engineer_why_we_go_from_clickhouse/",
		Year:    2023,
		From: []string{
			"clickhouse",
		},
		To: []string{
			"apachedoris",
		},
	},
	{
		Company: "Ongage",
		Uri:     "https://clickhouse.com/blog/ongages-strategic-shift-to-clickhouse-for-real-time-email-marketing",
		Year:    2023,
		From: []string{
			"mysql",
		},
		To: []string{
			"clickhouse",
		},
	},
	{
		Company: "Digger",
		Uri:     "https://medium.com/@DiggerHQ/we-rewrote-our-product-in-go-from-scratch-a0e9cbba193e",
		Year:    2023,
		From: []string{
			"python",
		},
		To: []string{
			"go",
		},
	},
	{
		Company: "eFishery",
		Uri:     "https://www.linkedin.com/feed/update/urn:li:activity:7148876278417215488?updateEntityUrn=urn%3Ali%3Afs_feedUpdate%3A%28V2%2Curn%3Ali%3Aactivity%3A7148876278417215488%29",
		Year:    2024,
		From: []string{
			"amazon_cloudfront",
		},
		To: []string{
			"cloudflare",
		},
	},
	{
		Company: "Vercel",
		Uri:     "https://www.youtube.com/watch?v=_SzvJJ3_6M0",
		Year:    2024,
		From: []string{
			"go",
		},
		To: []string{
			"rust",
		},
	},
	{
		Company: "Ramp",
		Uri:     "https://www.youtube.com/watch?v=7BtUgUb4gCs",
		Year:    2024,
		From: []string{
			"postgresql",
		},
		To: []string{
			"clickhouse",
		},
	},
	{
		Company: "Infiscal",
		Uri:     "https://infisical.com/blog/postgresql-migration-technical",
		Year:    2024,
		From: []string{
			"mongodb",
		},
		To: []string{
			"postgresql",
		},
	},
	{
		Company: "Reddit",
		Uri:     "https://www.infoq.com/news/2024/03/reddit-metadata-s3-postgres/",
		Year:    2024,
		From: []string{
			"s3_metadata",
		},
		To: []string{
			"aurora_postgresql",
		},
	},
	{
		Company: "Coralogix",
		Uri:     "https://thenewstack.io/from-postgres-to-scylladb-nosql-with-a-349x-speed-boost/",
		Year:    2024,
		From: []string{
			"postgresql",
		},
		To: []string{
			"scylladb",
		},
	},
	{
		Company: "Fablehenge",
		Uri:     "https://dusty.phillips.codes/2024/03/20/we-rewrote-our-react-app-in-svelte-in-three-weeks/",
		Year:    2024,
		From: []string{
			"react",
		},
		To: []string{
			"svelte",
		},
	},
	{
		Company: "Corsearch",
		Uri:     "https://www.youtube.com/watch?v=BuS8jFL9cvw",
		Year:    2024,
		From: []string{
			"mysql",
		},
		To: []string{
			"clickhouse",
		},
	},
	{
		Company: "Uber",
		Uri:     "https://www.infoq.com/news/2024/05/uber-dynamodb-ledgerstore/",
		Year:    2024,
		From: []string{
			"dynamodb",
		},
		To: []string{
			"ledgerstore",
		},
	},
	{
		Company: "37signals",
		Uri:     "https://world.hey.com/dhh/linux-as-the-new-developer-default-at-37signals-ef0823b7",
		Year:    2024,
		From: []string{
			"mac",
		},
		To: []string{
			"linux",
		},
	},
	{
		Company: "Beehiiv",
		Uri:     "https://www.youtube.com/watch?v=q44bQQoYOtY",
		Year:    2024,
		From: []string{
			"postgresql",
		},
		To: []string{
			"clickhouse",
		},
	},
	{
		Company: "Bonree",
		Uri:     "https://clickhouse.com/blog/bonree-replaces-zookeeper-with-clickhouse-keeper-for-drastically-improved-performance-and-reduced-costs",
		Year:    2024,
		From: []string{
			"zookeeper",
		},
		To: []string{
			"clickhousekeeper",
		},
	},
	{
		Company: "CommonRoom",
		Uri:     "https://clickhouse.com/blog/clickhouse-replaces-postgres-to-power-real-time-analytics-in-common-room-customer-portal",
		Year:    2024,
		From: []string{
			"postgresql",
		},
		To: []string{
			"clickhouse",
		},
	},
	{
		Company: "Shopee",
		Uri:     "https://clickhouse.com/blog/seeing-the-big-picture-shopees-journey-to-distributed-tracing-with-clickhouse",
		Year:    2024,
		From: []string{
			"druid",
			"hive",
			"presto",
		},
		To: []string{
			"clickhouse",
		},
	},
	{
		Company: "OpenMeter",
		Uri:     "https://clickhouse.com/blog/openmeter-real-time-usage-based-billing-powered-by-clickhouse-cloud",
		Year:    2024,
		From: []string{
			"postgresql",
			"kafka",
			"ksqldb",
		},
		To: []string{
			"clickhouse",
		},
	},
	{
		Company: "trip.com",
		Uri:     "https://clickhouse.com/blog/how-trip.com-migrated-from-elasticsearch-and-built-a-50pb-logging-solution-with-clickhouse",
		Year:    2024,
		From: []string{
			"elasticsearch",
		},
		To: []string{
			"clickhouse",
		},
	},
	{
		Company: "Didi",
		Uri:     "https://clickhouse.com/blog/didi-migrates-from-elasticsearch-to-clickHouse-for-a-new-generation-log-storage-system",
		Year:    2024,
		From: []string{
			"elasticsearch",
		},
		To: []string{
			"clickhouse",
		},
	},
	{
		Company: "Figma",
		Uri:     "https://www.figma.com/blog/figmas-journey-to-typescript-compiling-away-our-custom-programming-language/",
		Year:    2024,
		From: []string{
			"skew",
		},
		To: []string{
			"typescript",
		},
	},
	{
		Company: "FlockSafety",
		Uri:     "https://www.youtube.com/watch?v=dN4yrzn8Td4",
		Year:    2025,
		From: []string{
			"redshift",
		},
		To: []string{
			"clickhouse",
		},
	},
	{
		Company: "EpicGames",
		Uri:     "https://thenewstack.io/why-teams-are-ditching-dynamodb/",
		Year:    2025,
		From: []string{
			"dynamodb",
		},
		To: []string{
			"scylladb",
		},
	},
	{
		Company: "DigitalTurbine",
		Uri:     "https://thenewstack.io/why-teams-are-ditching-dynamodb/",
		Year:    2025,
		From: []string{
			"dynamodb",
		},
		To: []string{
			"scylladb",
		},
	},
	{
		Company: "DaggerIO",
		Uri:     "https://dagger.io/blog/replaced-react-with-go",
		Year:    2025,
		From: []string{
			"react",
		},
		To: []string{
			"go",
			"wasm",
		},
	},
	// Added by the maintainer
	{
		Company: "PayU",
		Uri:     "https://medium.com/payu-engineering/seamless-vault-migration-moving-from-consul-to-raft-with-minimal-downtime-54cfbc15ee3c",
		Year:    2025,
		From: []string{
			"consul",
		},
		To: []string{
			"raft",
		},
	},
	{
		Company: "Honeybadger",
		Uri:     "https://www.honeybadger.io/blog/sidekiq-to-karafka",
		Year:    2025,
		From: []string{
			"sidekiq",
		},
		To: []string{
			"kafka",
		},
	},
	{
		Company: "Tractian",
		Uri:     "https://www.scylladb.com/2025/03/04/tractian-migrate-mongodb-scylladb/",
		Year:    2025,
		From: []string{
			"mongodb",
		},
		To: []string{
			"scylladb",
		},
	},
	{
		Company: "Slice",
		Uri:     "https://engineering.sliceit.com/seamlessly-migrating-from-mongodb-to-postgresql-leveraging-data-platform-for-scalable-success-404e2577a47b",
		Year:    2025,
		From: []string{
			"mongodb",
		},
		To: []string{
			"postgresql",
		},
	},
	{
		Company: "AVEQ",
		Uri:     "https://aveq.info/modernizing-our-video-analytics-stack-from-rails-to-node-js",
		Year:    2025,
		From: []string{
			"ruby",
			"rubyonrails",
		},
		To: []string{
			"typescript",
			"express",
		},
	},
	{
		Company: "Cloudflare Developer Documentation",
		Uri:     "https://blog.cloudflare.com/open-source-all-the-way-down-upgrading-our-developer-documentation",
		Year:    2025,
		From: []string{
			"go",
			"hugo",
		},
		To: []string{
			"typescript",
			"astro",
		},
	},
	{
		Company: "Aquis Stock Exchange",
		Uri:     "https://aws.amazon.com/blogs/industries/how-aquis-created-a-cloud-native-exchange-to-meet-the-highest-regulatory-requirements-on-aws/",
		Year:    2025,
		From: []string{
			"self-managed",
		},
		To: []string{
			"cloud",
			"aws",
		},
	},
	{
		Company: "Checkly",
		Uri:     "https://www.checklyhq.com/blog/heroku-to-aws-migration",
		Year:    2025,
		From: []string{
			"heroku",
		},
		To: []string{
			"aws",
		},
	},
}

type migration struct {
	Company string
	Uri     string
	Year    int64
	From    []string
	To      []string
}
