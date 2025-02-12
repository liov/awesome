import akshare as ak
from pyecharts import options as opts
from pyecharts.charts import Line
from pyecharts.commons.utils import JsCode
from datetime import datetime, timedelta

# 设置茅台股票代码
symbol = "512690"

# 获取当前日期
end_date = datetime.now().strftime('%Y%m%d')

# 计算三年前的日期
start_date = (datetime.now() - timedelta(days=3 * 365)).strftime('%Y%m%d')

# 使用yfinance获取股票数据
data = ak.fund_etf_hist_em(symbol, period="daily", start_date=start_date, end_date=end_date, adjust="hfq")
data.set_index("日期", inplace=True)

# 提取数据中的日期和收盘价
dates = data.index.strftime('%Y-%m-%d').tolist()
closing_prices = data['收盘'].tolist()

# 创建 Line 图表
line_chart = Line()
line_chart.add_xaxis(xaxis_data=dates)
line_chart.add_yaxis(series_name="酒ETF股价走势",
                     y_axis=closing_prices,
                     markline_opts=opts.MarkLineOpts(
                         data=[opts.MarkLineItem(type_="average", name="平均值")]
                     )
                     )
line_chart.set_global_opts(
    title_opts=opts.TitleOpts(title="酒ETF股价走势图（近三年）"),
    xaxis_opts=opts.AxisOpts(type_="category"),
    yaxis_opts=opts.AxisOpts(is_scale=True),
    datazoom_opts=[
        opts.DataZoomOpts(
            pos_bottom="-2%",
            range_start=0,
            range_end=100,
            type_="inside"
        ),
        opts.DataZoomOpts(
            pos_bottom="-2%",
            range_start=0,
            range_end=100,
            type_="slider",
        ),
    ],
    toolbox_opts=opts.ToolboxOpts(
        feature={
            "dataZoom": {"yAxisIndex": "none"},
            "restore": {},
            "saveAsImage": {},
        }
    ),
)

# 渲染图表
line_chart.render("maotai_stock_trend_chart2.html")