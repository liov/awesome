import akshare as ak
import pandas as pd
from datetime import datetime, timedelta
import matplotlib.pyplot as plt
plt.rcParams['font.sans-serif'] = ['SimHei']  # 使用黑体

# 获取股票数据
symbol = "512690"
# 计算三年前的日期
start_date = (datetime.now() - timedelta(days=3 * 365)).strftime('%Y%m%d')
end_date = datetime.now().strftime('%Y%m%d')

data = ak.fund_etf_hist_em(symbol, period="daily", start_date=start_date, end_date=end_date, adjust="hfq")
print(data.columns)
# 处理字段命名，以符合 Backtrader 的要求
data.columns = [
    'date',
    'open',
    'close',
    'high',
    'low',
    'volume',
    "amount",
    "amplitude",
    "pctChg",
    "pctPrice",
    "turnover"
]
# 把 date 作为日期索引，以符合 Backtrader 的要求
data.index = pd.to_datetime(data['date'])
# 初始化交叉信号列
data['Signal'] = 0

# 计算每日收益率
data['Daily_Return'] = data['close'].pct_change()

# 计算策略信号
data['Signal'] = 0
data.loc[data['Daily_Return'] > 0, 'Signal'] = 1  # 以涨幅为信号，可根据需要修改条件

# 计算策略收益
data['Strategy_Return'] = data['Signal'].shift(1) * data['Daily_Return']

# 计算累计收益
data['Cumulative_Return'] = (1 + data['Strategy_Return']).cumprod()

# 绘制累计收益曲线
plt.figure(figsize=(10, 6))
plt.plot(data['Cumulative_Return'], label='Strategy Cumulative Return', color='b')
plt.plot(data['close'] / data['close'].iloc[0], label='Stock Cumulative Return', color='g')
plt.title("Cumulative Return of Strategy vs. Stock")
plt.xlabel("Date")
plt.ylabel("Cumulative Return")
plt.legend()
plt.show()