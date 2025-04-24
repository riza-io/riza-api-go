# Changelog

## 0.11.0 (2025-04-24)

Full Changelog: [v0.10.0...v0.11.0](https://github.com/riza-io/riza-api-go/compare/v0.10.0...v0.11.0)

### Features

* **api:** api update ([6c47882](https://github.com/riza-io/riza-api-go/commit/6c4788247485fff145a04b46f4abc1dcfedbf15c))
* **api:** api update ([81a6537](https://github.com/riza-io/riza-api-go/commit/81a653722fd4b276f18e3b4aced956c16f2145d8))
* **api:** api update ([c62d761](https://github.com/riza-io/riza-api-go/commit/c62d7614f047d4f1caddb9d1cf398bc5c28f54e7))
* **api:** api update ([6091cc8](https://github.com/riza-io/riza-api-go/commit/6091cc8b3800e85522f701051fba6a0f5e8b266a))
* **api:** api update ([#113](https://github.com/riza-io/riza-api-go/issues/113)) ([362e34c](https://github.com/riza-io/riza-api-go/commit/362e34c83912b929f013c6c225a4879a82327808))
* **client:** add support for reading base URL from environment variable ([2386a6f](https://github.com/riza-io/riza-api-go/commit/2386a6f50c8cb637ab5f758ab78183110ee20397))
* **client:** support custom http clients ([#114](https://github.com/riza-io/riza-api-go/issues/114)) ([f9ae9a9](https://github.com/riza-io/riza-api-go/commit/f9ae9a9989cd54566978723c95c9499458971e10))


### Bug Fixes

* **client:** return error on bad custom url instead of panic ([#112](https://github.com/riza-io/riza-api-go/issues/112)) ([279608f](https://github.com/riza-io/riza-api-go/commit/279608f36941947839cc427d32625974d46d72de))
* **test:** return early after test failure ([#110](https://github.com/riza-io/riza-api-go/issues/110)) ([f0f79cc](https://github.com/riza-io/riza-api-go/commit/f0f79cc384843928887446828494136331ca3d2f))


### Chores

* add request options to client tests ([#109](https://github.com/riza-io/riza-api-go/issues/109)) ([8ad7670](https://github.com/riza-io/riza-api-go/commit/8ad7670c0f2f708b458623737e7d181fb3d76a31))
* **ci:** add timeout thresholds for CI jobs ([afaa476](https://github.com/riza-io/riza-api-go/commit/afaa47664ca519bbcfeeece3f0b066d8751e0ac6))
* **ci:** only use depot for staging repos ([e5353eb](https://github.com/riza-io/riza-api-go/commit/e5353eb4a81747b945cf698ed831bd437f579bd7))
* **docs:** document pre-request options ([f95f41e](https://github.com/riza-io/riza-api-go/commit/f95f41e8dad9a091fb782e8274ee8408b4a0b627))
* **docs:** improve security documentation ([#107](https://github.com/riza-io/riza-api-go/issues/107)) ([2389440](https://github.com/riza-io/riza-api-go/commit/2389440489e6977ee71d57a67dcb74d7069e6e04))
* fix typos ([#111](https://github.com/riza-io/riza-api-go/issues/111)) ([4709229](https://github.com/riza-io/riza-api-go/commit/4709229aad1953fd0380d42a19cdf6f52ad4446c))
* **internal:** codegen related update ([2653cff](https://github.com/riza-io/riza-api-go/commit/2653cffd013c42787e2f45987605979b02178e30))
* **internal:** expand CI branch coverage ([d6f56db](https://github.com/riza-io/riza-api-go/commit/d6f56db1edb91ac8db14fb39f174d40ee51d732a))
* **internal:** reduce CI branch coverage ([a6a23ca](https://github.com/riza-io/riza-api-go/commit/a6a23caf08ce53d3756f01952f4be7f5963abeb2))


### Documentation

* update documentation links to be more uniform ([345e83f](https://github.com/riza-io/riza-api-go/commit/345e83fcd7763074c98211f55d6f38299a50706d))

## 0.10.0 (2025-03-14)

Full Changelog: [v0.9.0...v0.10.0](https://github.com/riza-io/riza-api-go/compare/v0.9.0...v0.10.0)

### Features

* **api:** api update ([#103](https://github.com/riza-io/riza-api-go/issues/103)) ([bf623a3](https://github.com/riza-io/riza-api-go/commit/bf623a30ad6fafbc243d76bd227971c72f434c31))
* **client:** improve default client options support ([#104](https://github.com/riza-io/riza-api-go/issues/104)) ([9078a6c](https://github.com/riza-io/riza-api-go/commit/9078a6cc3e7900b93323a644a79e84ac2fc918fc))


### Chores

* **internal:** codegen related update ([#101](https://github.com/riza-io/riza-api-go/issues/101)) ([8c118f3](https://github.com/riza-io/riza-api-go/commit/8c118f3964fae0529097496cb2ad970d39e591f8))
* **internal:** remove extra empty newlines ([#105](https://github.com/riza-io/riza-api-go/issues/105)) ([6fcfd1b](https://github.com/riza-io/riza-api-go/commit/6fcfd1b8cd6008b7f5edba3c211fcdf13fa27e16))

## 0.9.0 (2025-03-11)

Full Changelog: [v0.8.0...v0.9.0](https://github.com/riza-io/riza-api-go/compare/v0.8.0...v0.9.0)

### Features

* add SKIP_BREW env var to ./scripts/bootstrap ([#98](https://github.com/riza-io/riza-api-go/issues/98)) ([2a2cc6d](https://github.com/riza-io/riza-api-go/commit/2a2cc6d44d515a7a56eef9fc34e2118d830f8146))
* **api:** api update ([#90](https://github.com/riza-io/riza-api-go/issues/90)) ([1e429b4](https://github.com/riza-io/riza-api-go/commit/1e429b4472db4ca0f476807a4c2f2f145c4d26b1))
* **api:** api update ([#93](https://github.com/riza-io/riza-api-go/issues/93)) ([c474776](https://github.com/riza-io/riza-api-go/commit/c474776cd50ce5412814203a12c94218aeb1a82f))
* **api:** api update ([#97](https://github.com/riza-io/riza-api-go/issues/97)) ([d5c1364](https://github.com/riza-io/riza-api-go/commit/d5c136421ad9755cb2fe275de67c58ffcdc3fec0))
* **client:** accept RFC6838 JSON content types ([#99](https://github.com/riza-io/riza-api-go/issues/99)) ([08662f6](https://github.com/riza-io/riza-api-go/commit/08662f6a7412aa1cd75689343216b52dfc7fafa5))
* **client:** allow custom baseurls without trailing slash ([#96](https://github.com/riza-io/riza-api-go/issues/96)) ([a27c784](https://github.com/riza-io/riza-api-go/commit/a27c784408aa973a7919d8c4290ccec10c423240))


### Bug Fixes

* **client:** don't truncate manually specified filenames ([#91](https://github.com/riza-io/riza-api-go/issues/91)) ([335ebd9](https://github.com/riza-io/riza-api-go/commit/335ebd95cb7ac7c178dd3b90af58eee7d8a4a5c6))
* do not call path.Base on ContentType ([#89](https://github.com/riza-io/riza-api-go/issues/89)) ([99fbac5](https://github.com/riza-io/riza-api-go/commit/99fbac59fe49ba37de1d8f1de24792edff298071))
* fix early cancel when RequestTimeout is provided for streaming requests ([#88](https://github.com/riza-io/riza-api-go/issues/88)) ([6b9ff08](https://github.com/riza-io/riza-api-go/commit/6b9ff0863a0efd10f1a0b2c59d4d67afa8b2914e))


### Chores

* add UnionUnmarshaler for responses that are interfaces ([#86](https://github.com/riza-io/riza-api-go/issues/86)) ([44dab94](https://github.com/riza-io/riza-api-go/commit/44dab941130bb59919e2fdbc713fb694606360e3))
* **internal:** fix devcontainers setup ([#92](https://github.com/riza-io/riza-api-go/issues/92)) ([a2fd525](https://github.com/riza-io/riza-api-go/commit/a2fd5252bb8665bcbaec7a9ea0bf14610de2b95f))
* simplify string literals ([#95](https://github.com/riza-io/riza-api-go/issues/95)) ([24640e4](https://github.com/riza-io/riza-api-go/commit/24640e46a8ab927967be3218663f7fbf263bd69a))


### Documentation

* update URLs from stainlessapi.com to stainless.com ([#94](https://github.com/riza-io/riza-api-go/issues/94)) ([03af88e](https://github.com/riza-io/riza-api-go/commit/03af88e3fd0bb14b7f14203874382e5d2d7f1768))

## 0.8.0 (2025-02-04)

Full Changelog: [v0.7.0...v0.8.0](https://github.com/riza-io/riza-api-go/compare/v0.7.0...v0.8.0)

### Features

* **api:** api update ([#73](https://github.com/riza-io/riza-api-go/issues/73)) ([45a357c](https://github.com/riza-io/riza-api-go/commit/45a357cac0c0337b8ff995306a8b313dcfb1b1d4))
* **api:** api update ([#75](https://github.com/riza-io/riza-api-go/issues/75)) ([feffe59](https://github.com/riza-io/riza-api-go/commit/feffe593826d2de7dd00943a1d808b15ec9af777))
* **api:** api update ([#76](https://github.com/riza-io/riza-api-go/issues/76)) ([bbfad5b](https://github.com/riza-io/riza-api-go/commit/bbfad5b9ea1ad9a3c145b1cfe899ada739a5e6f0))
* **api:** api update ([#77](https://github.com/riza-io/riza-api-go/issues/77)) ([8ab051f](https://github.com/riza-io/riza-api-go/commit/8ab051fddbf908cf482c8945ab8ae2e15c459c12))
* **api:** api update ([#78](https://github.com/riza-io/riza-api-go/issues/78)) ([0525154](https://github.com/riza-io/riza-api-go/commit/0525154baa791111d4fa6b1b84919e79c208e228))
* **api:** api update ([#79](https://github.com/riza-io/riza-api-go/issues/79)) ([87ae631](https://github.com/riza-io/riza-api-go/commit/87ae631c476db0656984dd3a9fa495804143842a))
* **api:** api update ([#80](https://github.com/riza-io/riza-api-go/issues/80)) ([0274b2c](https://github.com/riza-io/riza-api-go/commit/0274b2cce301f580cd9d3a626dfed92b98470785))
* **client:** send `X-Stainless-Timeout` header ([#84](https://github.com/riza-io/riza-api-go/issues/84)) ([a587ed8](https://github.com/riza-io/riza-api-go/commit/a587ed85e2675f8b5c6a442131bf9e3dfb815eb0))


### Bug Fixes

* fix unicode encoding for json ([#82](https://github.com/riza-io/riza-api-go/issues/82)) ([1e7d130](https://github.com/riza-io/riza-api-go/commit/1e7d1302645c5863b77c1413add7123c5d0726db))


### Chores

* refactor client tests ([#81](https://github.com/riza-io/riza-api-go/issues/81)) ([111bb9b](https://github.com/riza-io/riza-api-go/commit/111bb9bb1072de946f756d904021b13ee72a17e0))


### Documentation

* document raw responses ([#83](https://github.com/riza-io/riza-api-go/issues/83)) ([32e915b](https://github.com/riza-io/riza-api-go/commit/32e915b0a00317cdd5850f74a6e7cf0b5e37b148))

## 0.7.0 (2025-01-23)

Full Changelog: [v0.6.0...v0.7.0](https://github.com/riza-io/riza-api-go/compare/v0.6.0...v0.7.0)

### Features

* **api:** api update ([#71](https://github.com/riza-io/riza-api-go/issues/71)) ([6ee7e40](https://github.com/riza-io/riza-api-go/commit/6ee7e409be7d31d82b196ebd19ad034343e55c8f))


### Chores

* **internal:** codegen related update ([#69](https://github.com/riza-io/riza-api-go/issues/69)) ([45d5b2c](https://github.com/riza-io/riza-api-go/commit/45d5b2c6adf566245bb3336cae9f6551ff3be5f0))

## 0.6.0 (2025-01-16)

Full Changelog: [v0.5.0...v0.6.0](https://github.com/riza-io/riza-api-go/compare/v0.5.0...v0.6.0)

### Features

* **api:** api update ([#66](https://github.com/riza-io/riza-api-go/issues/66)) ([81ab03e](https://github.com/riza-io/riza-api-go/commit/81ab03e250dd9279a759c4686b6985daa1b7a85f))

## 0.5.0 (2025-01-09)

Full Changelog: [v0.4.0...v0.5.0](https://github.com/riza-io/riza-api-go/compare/v0.4.0...v0.5.0)

### Features

* **api:** api update ([#61](https://github.com/riza-io/riza-api-go/issues/61)) ([aaef93d](https://github.com/riza-io/riza-api-go/commit/aaef93dc5ba0e981369c727be47185a9cbd8d1ad))
* **api:** api update ([#63](https://github.com/riza-io/riza-api-go/issues/63)) ([086af26](https://github.com/riza-io/riza-api-go/commit/086af26e109f8f33513ea2a673756c6a9f23c0f9))
* **api:** api update ([#64](https://github.com/riza-io/riza-api-go/issues/64)) ([4a05038](https://github.com/riza-io/riza-api-go/commit/4a050386b642467633ae8f6b75570a53c3249f73))


### Chores

* **internal:** codegen related update ([#58](https://github.com/riza-io/riza-api-go/issues/58)) ([67c178d](https://github.com/riza-io/riza-api-go/commit/67c178db0dfe3d38d936a2c6ded06b2febfa7419))
* **internal:** codegen related update ([#60](https://github.com/riza-io/riza-api-go/issues/60)) ([b11bd45](https://github.com/riza-io/riza-api-go/commit/b11bd45f65c7ceb74125c10227555324a0a6c4a0))
* **internal:** codegen related update ([#62](https://github.com/riza-io/riza-api-go/issues/62)) ([4ac1bf5](https://github.com/riza-io/riza-api-go/commit/4ac1bf55512a25b5a37da479bfc61d5325e6282d))

## 0.4.0 (2024-12-18)

Full Changelog: [v0.3.0...v0.4.0](https://github.com/riza-io/riza-api-go/compare/v0.3.0...v0.4.0)

### Features

* **api:** api update ([#53](https://github.com/riza-io/riza-api-go/issues/53)) ([1b77e01](https://github.com/riza-io/riza-api-go/commit/1b77e01a8e8b6a0ae669e75dd791a70b39d29b8e))
* **api:** api update ([#55](https://github.com/riza-io/riza-api-go/issues/55)) ([bd82cd5](https://github.com/riza-io/riza-api-go/commit/bd82cd5692422a0ff5db0717601158bbef185069))
* **api:** api update ([#56](https://github.com/riza-io/riza-api-go/issues/56)) ([1cfbe6e](https://github.com/riza-io/riza-api-go/commit/1cfbe6eac34a9bc6e2132ab5407ee8d07b4579f3))

## 0.3.0 (2024-12-02)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/riza-io/riza-api-go/compare/v0.2.0...v0.3.0)

### Features

* **api:** api update ([#49](https://github.com/riza-io/riza-api-go/issues/49)) ([f210c04](https://github.com/riza-io/riza-api-go/commit/f210c047c22fce6a9c7baff3c6e818e50a685544))
* **api:** api update ([#51](https://github.com/riza-io/riza-api-go/issues/51)) ([74c2f77](https://github.com/riza-io/riza-api-go/commit/74c2f77f0fc3aa80de1da6001a3638736a09d5fc))


### Chores

* rebuild project due to codegen change ([#46](https://github.com/riza-io/riza-api-go/issues/46)) ([8bd8d40](https://github.com/riza-io/riza-api-go/commit/8bd8d40bf2491e1c5a01abb56c85e7cfb7525d1c))
* rebuild project due to codegen change ([#48](https://github.com/riza-io/riza-api-go/issues/48)) ([661b38a](https://github.com/riza-io/riza-api-go/commit/661b38a6511e84da23d0d471da0ee008b17127d1))
* rebuild project due to codegen change ([#50](https://github.com/riza-io/riza-api-go/issues/50)) ([80e0064](https://github.com/riza-io/riza-api-go/commit/80e00642d28150bafd7cc5cfa46a65251a6e7ab5))

## 0.2.0 (2024-11-07)

Full Changelog: [v0.1.1...v0.2.0](https://github.com/riza-io/riza-api-go/compare/v0.1.1...v0.2.0)

### Features

* **api:** api update ([#41](https://github.com/riza-io/riza-api-go/issues/41)) ([274d3e3](https://github.com/riza-io/riza-api-go/commit/274d3e3541722d68f3da1019bb8d58cdbc54b4a2))
* **api:** api update ([#44](https://github.com/riza-io/riza-api-go/issues/44)) ([220cc62](https://github.com/riza-io/riza-api-go/commit/220cc62f8c817a203720415c58e5e10978c8c4f8))
* **api:** manual updates ([#43](https://github.com/riza-io/riza-api-go/issues/43)) ([465baad](https://github.com/riza-io/riza-api-go/commit/465baad08d14d0cf6c79100e97d7f135b2271d5e))

## 0.1.1 (2024-11-06)

Full Changelog: [v0.1.0-alpha.9...v0.1.1](https://github.com/riza-io/riza-api-go/compare/v0.1.0-alpha.9...v0.1.1)

### Features

* **api:** api update ([#34](https://github.com/riza-io/riza-api-go/issues/34)) ([c7a81b3](https://github.com/riza-io/riza-api-go/commit/c7a81b3794fb4ed4a6be49cc2e303e4ee105cd0b))
* **api:** api update ([#36](https://github.com/riza-io/riza-api-go/issues/36)) ([58933d4](https://github.com/riza-io/riza-api-go/commit/58933d4d3f6bcaf8994b3beeeab8392d95db6720))
* **api:** api update ([#37](https://github.com/riza-io/riza-api-go/issues/37)) ([32a53f7](https://github.com/riza-io/riza-api-go/commit/32a53f79bb27cab5ec5d57b19eba0badfb7e6e56))
* **api:** api update ([#38](https://github.com/riza-io/riza-api-go/issues/38)) ([7037dea](https://github.com/riza-io/riza-api-go/commit/7037dead76084b37bab188d4746d9d17d95e90b5))
* **api:** api update ([#39](https://github.com/riza-io/riza-api-go/issues/39)) ([aee7056](https://github.com/riza-io/riza-api-go/commit/aee705620a5328ef30d1856ded510663819e6f03))

## 0.1.0-alpha.9 (2024-10-24)

Full Changelog: [v0.1.0-alpha.8...v0.1.0-alpha.9](https://github.com/riza-io/riza-api-go/compare/v0.1.0-alpha.8...v0.1.0-alpha.9)

### Features

* **api:** api update ([#30](https://github.com/riza-io/riza-api-go/issues/30)) ([4249a5a](https://github.com/riza-io/riza-api-go/commit/4249a5a35496499320cf5993c63eac438352d6bc))
* **api:** manual updates ([#32](https://github.com/riza-io/riza-api-go/issues/32)) ([d476d85](https://github.com/riza-io/riza-api-go/commit/d476d8547b234690dbd7b8f37686ac68e175e709))

## 0.1.0-alpha.8 (2024-09-13)

Full Changelog: [v0.1.0-alpha.7...v0.1.0-alpha.8](https://github.com/riza-io/riza-api-go/compare/v0.1.0-alpha.7...v0.1.0-alpha.8)

### Features

* **api:** OpenAPI spec update via Stainless API ([#23](https://github.com/riza-io/riza-api-go/issues/23)) ([3eeaab1](https://github.com/riza-io/riza-api-go/commit/3eeaab1625a9724ab26d132477b1941677db8017))


### Bug Fixes

* **requestconfig:** copy over more fields when cloning ([#26](https://github.com/riza-io/riza-api-go/issues/26)) ([b4f4a13](https://github.com/riza-io/riza-api-go/commit/b4f4a13898477e4f8bb9d78ba212b114917d4df8))


### Chores

* **internal:** codegen related update ([#24](https://github.com/riza-io/riza-api-go/issues/24)) ([c8f6136](https://github.com/riza-io/riza-api-go/commit/c8f6136fabff019582c1e83bcae4c6e4140d4642))
* **internal:** codegen related update ([#27](https://github.com/riza-io/riza-api-go/issues/27)) ([c8ba87a](https://github.com/riza-io/riza-api-go/commit/c8ba87a68699b7dc49835b4f628d41247db0ee95))


### Documentation

* update CONTRIBUTING.md ([#28](https://github.com/riza-io/riza-api-go/issues/28)) ([27e70ff](https://github.com/riza-io/riza-api-go/commit/27e70ffe7f1553404478edf5be322a7588d2c7b2))

## 0.1.0-alpha.7 (2024-07-23)

Full Changelog: [v0.1.0-alpha.6...v0.1.0-alpha.7](https://github.com/riza-io/riza-api-go/compare/v0.1.0-alpha.6...v0.1.0-alpha.7)

### Features

* **api:** OpenAPI spec update via Stainless API ([#17](https://github.com/riza-io/riza-api-go/issues/17)) ([eb86e3d](https://github.com/riza-io/riza-api-go/commit/eb86e3db6249bf72e1fe313c8fcc02d85b264f1a))


### Chores

* **ci:** limit release doctor target branches ([#20](https://github.com/riza-io/riza-api-go/issues/20)) ([674c19b](https://github.com/riza-io/riza-api-go/commit/674c19b631604a469fa9829badfed40313dbc9e7))
* **ci:** remove unused release doctor ([#21](https://github.com/riza-io/riza-api-go/issues/21)) ([6eadd74](https://github.com/riza-io/riza-api-go/commit/6eadd7428bfeb64239bb57cdea3d0a9fb240b5a2))
* **internal:** codegen related update ([#18](https://github.com/riza-io/riza-api-go/issues/18)) ([b7e3a7d](https://github.com/riza-io/riza-api-go/commit/b7e3a7d073e3ce95c68ca5cf2c2556ba7fd9ae55))
* **tests:** update prism version ([#22](https://github.com/riza-io/riza-api-go/issues/22)) ([ab187b8](https://github.com/riza-io/riza-api-go/commit/ab187b8b4cf37d0fdc24e6e58366785c9f40659e))

## 0.1.0-alpha.6 (2024-05-22)

Full Changelog: [v0.1.0-alpha.5...v0.1.0-alpha.6](https://github.com/riza-io/riza-api-go/compare/v0.1.0-alpha.5...v0.1.0-alpha.6)

### Features

* **api:** OpenAPI spec update via Stainless API ([#13](https://github.com/riza-io/riza-api-go/issues/13)) ([ee07454](https://github.com/riza-io/riza-api-go/commit/ee07454b71782f7a5a05528ed2d453721edc2d71))
* **api:** OpenAPI spec update via Stainless API ([#15](https://github.com/riza-io/riza-api-go/issues/15)) ([b018f31](https://github.com/riza-io/riza-api-go/commit/b018f31bf878771e88bab46893854c994a6d969e))

## 0.1.0-alpha.5 (2024-04-23)

Full Changelog: [v0.1.0-alpha.4...v0.1.0-alpha.5](https://github.com/riza-io/riza-api-go/compare/v0.1.0-alpha.4...v0.1.0-alpha.5)

### Features

* **api:** update via SDK Studio ([#11](https://github.com/riza-io/riza-api-go/issues/11)) ([f3e5c96](https://github.com/riza-io/riza-api-go/commit/f3e5c961a994cfce9f07f1b18945227becd649ca))

## 0.1.0-alpha.4 (2024-04-19)

Full Changelog: [v0.1.0-alpha.3...v0.1.0-alpha.4](https://github.com/riza-io/riza-api-go/compare/v0.1.0-alpha.3...v0.1.0-alpha.4)

### Features

* **api:** OpenAPI spec update via Stainless API ([#9](https://github.com/riza-io/riza-api-go/issues/9)) ([72fd0bd](https://github.com/riza-io/riza-api-go/commit/72fd0bd5e1307aeef016a7d0492586bc02062a5c))

## 0.1.0-alpha.3 (2024-04-18)

Full Changelog: [v0.1.0-alpha.2...v0.1.0-alpha.3](https://github.com/riza-io/riza-api-go/compare/v0.1.0-alpha.2...v0.1.0-alpha.3)

### Features

* **api:** update via SDK Studio ([#7](https://github.com/riza-io/riza-api-go/issues/7)) ([abb3637](https://github.com/riza-io/riza-api-go/commit/abb3637374d13cab4eb29fe21d748ef9f8b4a641))

## 0.1.0-alpha.2 (2024-04-18)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha.2](https://github.com/riza-io/riza-api-go/compare/v0.1.0-alpha.1...v0.1.0-alpha.2)

### Features

* **api:** update via SDK Studio ([#3](https://github.com/riza-io/riza-api-go/issues/3)) ([c05ddf6](https://github.com/riza-io/riza-api-go/commit/c05ddf6d8d797ac9c497869b01bf4e18f516c238))
* **api:** update via SDK Studio ([#5](https://github.com/riza-io/riza-api-go/issues/5)) ([7592e52](https://github.com/riza-io/riza-api-go/commit/7592e52ba20def022458699f0fb7791ce436bb4e))
* **api:** update via SDK Studio ([#6](https://github.com/riza-io/riza-api-go/issues/6)) ([9cd1e86](https://github.com/riza-io/riza-api-go/commit/9cd1e86d40bd6651ca6c383598445ad50543e80f))

## 0.1.0-alpha.1 (2024-04-17)

Full Changelog: [v0.0.1-alpha.0...v0.1.0-alpha.1](https://github.com/riza-io/riza-api-go/compare/v0.0.1-alpha.0...v0.1.0-alpha.1)

### Features

* **api:** update via SDK Studio ([1313ca6](https://github.com/riza-io/riza-api-go/commit/1313ca673b6668a3e9b5f65785251145d4e6b0dc))
* **api:** update via SDK Studio ([5032cc0](https://github.com/riza-io/riza-api-go/commit/5032cc03df948438e0ed8201cd5b11d2ce7e7940))


### Chores

* configure new SDK language ([e90b3c5](https://github.com/riza-io/riza-api-go/commit/e90b3c5c9e49ec8df1385de14781fee1f0e0f061))
* go live ([#1](https://github.com/riza-io/riza-api-go/issues/1)) ([7edc050](https://github.com/riza-io/riza-api-go/commit/7edc0507a223e98e8fa0be39a21d4453a53e3e44))
