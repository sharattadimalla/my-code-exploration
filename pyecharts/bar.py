from pyecharts.charts import Bar
from pyecharts import options as opts

b = Bar()

b.add_xaxis(["M", "A", "I", "O", "G", "Al"])

b.add_yaxis("Revenue $", [70,100,50,90,60])

b.set_global_opts(title_opts=opts.TitleOpts(title="Title", subtitle="Revenue"),
                     toolbox_opts=opts.ToolboxOpts())

b.render()
