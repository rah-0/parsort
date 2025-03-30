# Tuner

As observed by Reddit user **LearnedByError** in [this comment](https://www.reddit.com/r/golang/comments/1jliue6/comment/mk94phj/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button), the performance with the library on `Intel(R) Xeon(R) E5-2680` is quite poor.

So poor in fact, that using the library might not even be worth it with old CPUs.

Nevertheless and out of curiosity I have created [Tuner](https://github.com/rah-0/parsort/blob/master/tuner.go#L534) which is a tool that helps benchmark and updates configuration thresholds per type. 

Tuning is optional, you can either simply call `parsort.Tune()` which uses default values or [parsort.TuneSpecific](https://github.com/rah-0/parsort/blob/master/tuner.go#L27) to have more control.

Changing any [Config](https://github.com/rah-0/parsort/blob/master/config.go)ParallelSize variable has an impact of when parallelization starts.

