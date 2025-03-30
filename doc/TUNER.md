# Tuner

As observed by Reddit user **LearnedByError** in [this comment](https://www.reddit.com/r/golang/comments/1jliue6/comment/mk94phj/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button), the performance with the library on `Intel(R) Xeon(R) E5-2680` is quite poor.

So poor in fact, that using the library might not even be worth it with old CPUs.

Nevertheless and out of curiosity I have created [Tuner](https://github.com/rah-0/parsort/blob/master/tuner.go#L534) which is a tool that helps benchmark and updates configuration thresholds per type. 

Tuning is optional, you can either simply call `parsort.Tune()` which uses default values or [parsort.TuneSpecific](https://github.com/rah-0/parsort/blob/master/tuner.go#L27) to have more control.

Changing any [Config](https://github.com/rah-0/parsort/blob/master/config.go)ParallelSize variable has an impact of when parallelization starts.

## Tuner output of systems

### AMD Ryzen 5950X 5GHz + 64GB DDR4 3600MHz CL19
| Type     | Label | Sequential ns/op | Concurrent ns/op | Sequential B/op | Concurrent B/op | Delta ns/op | Delta B/op |
|----------|-------|------------------|------------------|------------------|------------------|-------------|-----------|
| Int      | 3000  | 119µs 97ns       | 105µs 331ns      | 24KB             | 149KB 979B       | -11.56%     | +524.82%  |
| Int8     | 3000  | 137µs 121ns      | 118µs 93ns       | 3KB 56B          | 25KB 726B        | -13.88%     | +741.62%  |
| Int16    | 2000  | 114µs 583ns      | 98µs 645ns       | 4KB 56B          | 31KB 634B        | -13.91%     | +679.82%  |
| Int32    | 2000  | 121µs 541ns      | 104µs 845ns      | 8KB 56B          | 55KB 626B        | -13.74%     | +590.42%  |
| Int64    | 2000  | 130µs 541ns      | 110µs 404ns      | 16KB 56B         | 103KB 571B       | -15.43%     | +545.03%  |
| Uint     | 3000  | 197µs 73ns       | 142µs 351ns      | 24KB 56B         | 151KB 694B       | -27.77%     | +530.55%  |
| Uint8    | 2000  | 96µs 63ns        | 85µs 803ns       | 2KB 56B          | 19KB 662B        | -10.68%     | +856.18%  |
| Uint16   | 2000  | 120µs 106ns      | 93µs 586ns       | 4KB 56B          | 31KB 631B        | -22.08%     | +679.74%  |
| Uint32   | 2000  | 129µs 846ns      | 108µs 297ns      | 8KB 56B          | 55KB 630B        | -16.60%     | +590.47%  |
| Uint64   | 3000  | 197µs 651ns      | 156µs 712ns      | 24KB 56B         | 151KB 697B       | -20.71%     | +530.57%  |
| Float32  | 2000  | 137µs 196ns      | 104µs 618ns      | 8KB 56B          | 55KB 632B        | -23.75%     | +590.49%  |
| Float64  | 2000  | 102µs 849ns      | 88µs 100ns       | 16KB             | 101KB 818B       | -14.34%     | +536.24%  |
| String   | 2000  | 170µs 972ns      | 137µs 256ns      | 32KB             | 197KB 699B       | -19.72%     | +517.76%  |
| Time     | 2000  | 243µs 177ns      | 175µs 181ns      | 48KB 96B         | 296KB 564B       | -27.96%     | +516.61%  |
| Struct   | 9000  | 386µs 464ns      | 343µs 648ns      | 216KB 24B        | 442KB 75B        | -11.08%     | +104.64%  |

### AMD Ryzen AI 9 HX 370 + 32GB LPDDR5X 7500MHz CL40
| Type    | Label | Sequential ns/op | Concurrent ns/op | Sequential B/op | Concurrent B/op | Delta ns/op | Delta B/op |
|---------|-------|------------------|------------------|------------------|------------------|-------------|------------|
| Int     | 4000  | 281µs 699ns      | 210µs 927ns      | 32KB             | 184KB 648B       | -25.12%     | +476.98%   |
| Int8    | 2000  | 162µs 994ns      | 141µs 420ns      | 2KB 56B          | 17KB 304B        | -13.24%     | +741.83%   |
| Int16   | 2000  | 177µs 787ns      | 143µs 355ns      | 4KB 56B          | 28KB 696B        | -19.37%     | +607.32%   |
| Int32   | 3000  | 321µs 114ns      | 164µs 845ns      | 12KB 56B         | 73KB 782B        | -48.66%     | +511.91%   |
| Int64   | 3000  | 332µs 148ns      | 242µs 518ns      | 24KB 56B         | 141KB 867B       | -26.98%     | +489.68%   |
| Uint    | 3000  | 341µs 174ns      | 206µs 797ns      | 24KB 56B         | 141KB 844B       | -39.39%     | +489.59%   |
| Uint8   | 3000  | 244µs 486ns      | 169µs 152ns      | 3KB 56B          | 22KB 808B        | -30.81%     | +646.04%   |
| Uint16  | 2000  | 196µs 52ns       | 144µs 786ns      | 4KB 56B          | 28KB 696B        | -26.15%     | +607.32%   |
| Uint32  | 3000  | 312µs 186ns      | 220µs 323ns      | 12KB 56B         | 73KB 799B        | -29.43%     | +512.05%   |
| Uint64  | 2000  | 220µs 473ns      | 163µs 498ns      | 16KB 56B         | 96KB 274B        | -25.84%     | +499.62%   |
| Float32 | 3000  | 365µs 292ns      | 278µs 639ns      | 12KB 56B         | 73KB 783B        | -23.72%     | +511.92%   |
| Float64 | 3000  | 270µs 978ns      | 186µs 395ns      | 24KB             | 140KB 464B       | -31.21%     | +485.22%   |
| String  | 2000  | 324µs 240ns      | 239µs 40ns       | 32KB             | 192KB 468B       | -26.28%     | +501.43%   |
| Time    | 1000  | 233µs 360ns      | 196µs 363ns      | 24KB 96B         | 142KB 732B       | -15.85%     | +492.33%   |
| Struct  | 10000 | 733µs 238ns      | 652µs 297ns      | 240KB 24B        | 487KB 702B       | -11.04%     | +103.18%   |
