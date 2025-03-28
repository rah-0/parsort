## Baseline Bench
All benchmarks are performed with: **1.000**, **10.000**, **100.000**, **1.000.000**, **10.000.000** values.  
The baseline benches the standard lib, the results are the following:

### Int Asc
```
Asc_Int_1000-8         	  100512	     12235 ns/op	    8192 B/op	       1 allocs/op
Asc_Int_10000-8        	    2882	    406320 ns/op	   81920 B/op	       1 allocs/op
Asc_Int_100000-8       	     224	   5358521 ns/op	  802821 B/op	       1 allocs/op
Asc_Int_1000000-8      	      19	  60016454 ns/op	 8003589 B/op	       1 allocs/op
Asc_Int_10000000-8     	       2	 686426480 ns/op	80003072 B/op	       1 allocs/op
```
### Int Desc
```
Desc_Int_1000-8         	   28135	     42971 ns/op	    8232 B/op	       3 allocs/op
Desc_Int_10000-8        	    1280	    925838 ns/op	   81960 B/op	       3 allocs/op
Desc_Int_100000-8       	      99	  11773261 ns/op	  802918 B/op	       3 allocs/op
Desc_Int_1000000-8      	       8	 137250452 ns/op	 8003636 B/op	       3 allocs/op
Desc_Int_10000000-8     	       1	1591991046 ns/op	80003112 B/op	       3 allocs/op
```

### Float Asc
```
Asc_Float_1000-8         	   79406	     15260 ns/op	    8192 B/op	       1 allocs/op
Asc_Float_10000-8        	    2112	    562764 ns/op	   81920 B/op	       1 allocs/op
Asc_Float_100000-8       	     164	   7243807 ns/op	  802818 B/op	       1 allocs/op
Asc_Float_1000000-8      	      13	  81742600 ns/op	 8003584 B/op	       1 allocs/op
Asc_Float_10000000-8     	       2	 946438351 ns/op	80003072 B/op	       1 allocs/op
```
### Float Desc
```
Desc_Float_1000-8         	   23719	     50173 ns/op	    8232 B/op	       3 allocs/op
Desc_Float_10000-8        	    1075	   1090956 ns/op	   81960 B/op	       3 allocs/op
Desc_Float_100000-8       	      85	  13821467 ns/op	  802859 B/op	       3 allocs/op
Desc_Float_1000000-8      	       7	 160525362 ns/op	 8003637 B/op	       3 allocs/op
Desc_Float_10000000-8     	       1	1892565552 ns/op	80003112 B/op	       3 allocs/op
```

### String Asc
```
Asc_String_1000-8         	   18469	     63118 ns/op	   16384 B/op	       1 allocs/op
Asc_String_10000-8        	    1078	   1079163 ns/op	  163840 B/op	       1 allocs/op
Asc_String_100000-8       	      78	  14687354 ns/op	 1605632 B/op	       1 allocs/op
Asc_String_1000000-8      	       5	 225329861 ns/op	16007168 B/op	       1 allocs/op
Asc_String_10000000-8     	       1	3365926285 ns/op	160006144 B/op	       1 allocs/op
```
### String Desc
```
Desc_String_1000-8         	   13168	     90850 ns/op	   16424 B/op	       3 allocs/op
Desc_String_10000-8        	     777	   1521293 ns/op	  163880 B/op	       3 allocs/op
Desc_String_100000-8       	      50	  21462943 ns/op	 1605672 B/op	       3 allocs/op
Desc_String_1000000-8      	       4	 316241439 ns/op	16008520 B/op	       4 allocs/op
Desc_String_10000000-8     	       1	5236124366 ns/op	160006184 B/op	       3 allocs/op
```

### time.Time Asc/Desc
```
Sort_Time_1000-8         	   10000	    104695 ns/op	   24672 B/op	       4 allocs/op
Sort_Time_10000-8        	     800	   1546552 ns/op	  245862 B/op	       4 allocs/op
Sort_Time_100000-8       	      62	  18585826 ns/op	 2400437 B/op	       4 allocs/op
Sort_Time_1000000-8      	       5	 204948782 ns/op	24002656 B/op	       4 allocs/op
Sort_Time_10000000-8     	       1	2331454407 ns/op	240001120 B/op	       4 allocs/op
```

## Parsort Bench

### Int Asc
```
Parsort_Asc_Int_10000-8        	    3169	    377494 ns/op	  329166 B/op	      47 allocs/op
Parsort_Asc_Int_100000-8       	     350	   3457374 ns/op	 3229117 B/op	      47 allocs/op
Parsort_Asc_Int_1000000-8      	      67	  16378208 ns/op	32048752 B/op	      47 allocs/op
Parsort_Asc_Int_10000000-8     	       6	 174460002 ns/op	320030152 B/op	      47 allocs/op
```
### Int Desc
```
Parsort_Desc_Int_10000-8        	    3189	    377779 ns/op	  329178 B/op	      47 allocs/op
Parsort_Desc_Int_100000-8       	     422	   3505281 ns/op	 3229122 B/op	      47 allocs/op
Parsort_Desc_Int_1000000-8      	      67	  16796855 ns/op	32048661 B/op	      47 allocs/op
Parsort_Desc_Int_10000000-8     	       6	 182859490 ns/op	320030376 B/op	      47 allocs/op
```

### Float64 Asc
```
Parsort_Asc_Float64_10000-8        	    2722	    433885 ns/op	  329169 B/op	      47 allocs/op
Parsort_Asc_Float64_100000-8       	     309	   3452451 ns/op	 3229160 B/op	      47 allocs/op
Parsort_Asc_Float64_1000000-8      	      52	  21521789 ns/op	32048636 B/op	      47 allocs/op
Parsort_Asc_Float64_10000000-8     	       5	 220576750 ns/op	320030136 B/op	      47 allocs/op
```
### Float64 Desc
```
Parsort_Desc_Float64_10000-8        	    2798	    440571 ns/op	  329175 B/op	      47 allocs/op
Parsort_Desc_Float64_100000-8       	     434	   3371673 ns/op	 3229121 B/op	      47 allocs/op
Parsort_Desc_Float64_1000000-8      	      55	  21498810 ns/op	32048576 B/op	      47 allocs/op
Parsort_Desc_Float64_10000000-8     	       5	 223893618 ns/op	320030174 B/op	      47 allocs/op
```

### String Asc
```
Parsort_Asc_String_10000-8        	    1376	    842687 ns/op	  656868 B/op	      47 allocs/op
Parsort_Asc_String_100000-8       	     140	   7522620 ns/op	 6424024 B/op	      47 allocs/op
Parsort_Asc_String_1000000-8      	      10	 108093949 ns/op	64046520 B/op	      47 allocs/op
Parsort_Asc_String_10000000-8     	       1	1419925603 ns/op	640026040 B/op	      47 allocs/op
```
### String Desc
```
Parsort_Desc_String_10000-8        	    1364	    851547 ns/op	  656874 B/op	      47 allocs/op
Parsort_Desc_String_100000-8       	     148	   7818514 ns/op	 6424022 B/op	      47 allocs/op
Parsort_Desc_String_1000000-8      	      10	 105319618 ns/op	64046520 B/op	      47 allocs/op
Parsort_Desc_String_10000000-8     	       1	1436109086 ns/op	640026136 B/op	      48 allocs/op
```

### Time Asc/Desc
```
Parsort_Sort_Time_10000-8        	    1210	    992026 ns/op	 1001720 B/op	      71 allocs/op
Parsort_Sort_Time_100000-8       	     170	   6785655 ns/op	 9636026 B/op	      71 allocs/op
Parsort_Sort_Time_1000000-8      	      18	  60761972 ns/op	96028856 B/op	      71 allocs/op
Parsort_Sort_Time_10000000-8     	       2	 651085630 ns/op	960039096 B/op	      71 allocs/op
```
