import akshare as ak
import pandas as pd
from datetime import datetime, timedelta
import matplotlib.pyplot as plt
plt.rcParams['font.sans-serif'] = ['SimHei']  # 使用黑体

# 获取股票数据
symbol = "512690"
start_date = (datetime.now() - timedelta(days=365)).strftime('%Y%m%d')
end_date = datetime.now().strftime('%Y%m%d')

data = ak.fund_etf_hist_em(symbol, period="daily", start_date=start_date, end_date=end_date, adjust="")
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
print(data)
# 绘制股价走势图
# data['收盘'].plot(figsize=(10, 6), label=symbol)
# plt.title(f"{symbol} Stock Price")
# plt.xlabel("日期")
# plt.ylabel("价格")
# plt.legend()
# plt.show()

# 计算短期（50天）和长期（200天）移动平均
data['MA_10'] = data['close'].rolling(window=10).mean()
data['MA_40'] = data['close'].rolling(window=40).mean()


# 生成买卖信号
data['Signal'] = 0
# 修改买入信号
data.loc[data['MA_10'] > data['MA_40'], 'Signal'] = 1

# 修改卖出信号
data.loc[data['MA_10'] < data['MA_40'], 'Signal'] = -1


# 绘制股价和移动平均线
plt.figure(figsize=(10, 6))
plt.plot(data['close'], label='Close Price')
plt.plot(data['MA_10'], label='50-day Moving Average')
plt.plot(data['MA_40'], label='200-day Moving Average')

# 标记买卖信号
plt.scatter(data[data['Signal'] == 1].index, data[data['Signal'] == 1]['MA_10'], marker='^', color='g', label='Buy Signal')
plt.scatter(data[data['Signal'] == -1].index, data[data['Signal'] == -1]['MA_10'], marker='v', color='r', label='Sell Signal')

plt.title("Maotai Stock Price with Moving Averages")
plt.xlabel("Date")
plt.ylabel("Price (CNY)")
plt.legend()
plt.show()