Pipe-Filter Example:

            SplitFilter                    TointFilter               SumFilter
"1, 2, 3" --------------> ["1", "2", "3"] -------------> [1, 2, 3] -------------> 6

1. 適合與數據處理及數據分析系統
2. Filter 封裝數據處理功能
3. 鬆偶合: Filter 只跟數據(格式)耦合
4. Pipe 用於連接 Filter 傳遞數據或是在意不處理過程中緩衝 Data Stream
  - Process 內同步時調用時，pipe 演變為數據在方法調用間傳遞