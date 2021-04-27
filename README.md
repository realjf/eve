# eve
Natural Language Processing in GO


以下是已实现的功能列表：

- 文本标记
- 分句
- 形态分析
- 后缀处理，代词的重新标记
- 灵活的多字识别
- 收缩分裂
- 未知词类的概率预测
- 命名实体检测
- 词性标注
- 基于图表的浅层分析
- 命名实体分类（带有外部库MITIE-https://github.com/mit-nlp/MITIE)
- 基于规则的依赖关系分析

### 构建
```sh
# 构建外部库MITIE
git clone https://github.com/mit-nlp/MITIE
cd MITIE
make
cp mitielib/libmitie.so /usr/lib/

# 构建本项目
go build eve.go
```


### 如何使用api
```sh
HTTP GET: http://localhost:9999/analyzer?url={COPY HERE AN URL}
```


```sh
HTTP POST:

http://localhost:9999/analyzer-api

{
    content: 'Text you want to analyze'
}
```