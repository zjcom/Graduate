/**
 * Created by piasy on 15/5/18.
 */

function createDistWithHRChart(players) {
    Highcharts.setOptions({
        global: {
            useUTC: false
        }
    });

    var player_names = [];
    players.forEach(function (p, i, arr) {
        player_names.push(p.name);
    });

    var series = [];
    $('#chart_container').highcharts({
        chart: {
            type: 'column',
            animation: Highcharts.svg, // don't animate in old IE
            marginRight: 10,
            events: {
                load: function() {
                    series = this.series;
                }
            }
        },
        title: {
            text: '不同心率跑动距离统计'
        },
        xAxis: {
            categories: player_names
        },
        yAxis: {
            min: 0,
            title: {
                text: '距离（米）'
            },
            stackLabels: {
                enabled: true,
                style: {
                    fontWeight: 'bold',
                    color: (Highcharts.theme && Highcharts.theme.textColor) || 'gray'
                }
            }
        },
        tooltip: {
            formatter: function() {
                return '<b>'+ this.x +'</b><br/>'+
                    this.series.name +': '+ this.y.toFixed(2) +'<br/>'+
                    '总距离：'+ this.point.stackTotal.toFixed(2);
            }
        },
        plotOptions: {
            column: {
                stacking: 'normal',
                dataLabels: {
                    enabled: true,
                    color: (Highcharts.theme && Highcharts.theme.dataLabelsColor) || 'white'
                }
            }
        },
        exporting: {
            enabled: false
        },
        credits: {
            enabled: false
        },
        series: [{
            name: '~ 120',
            data: []
        }, {
            name: '120 ~ 150',
            data: []
        }, {
            name: '150 ~',
            data: []
        }]
    });
    return series;
}

function updateDistWithHRChartFirstTime(series, players) {
    var data1 = [];
    var data2 = [];
    var data3 = [];
    players.forEach(function (p, i, arr) {
        var remark = JSON.parse(localStorage.getItem(p.history));
        if (remark != null && remark.distwithhr != null && remark.distwithhr.length == 3) {
            data1.push(Number(remark.distwithhr[0].toFixed(2)));
            data2.push(Number(remark.distwithhr[1].toFixed(2)));
            data3.push(Number(remark.distwithhr[2].toFixed(2)));
        } else {
            data1.push(0);
            data2.push(0);
            data3.push(0);
        }
    });
    series[0].setData(data1);
    series[1].setData(data2);
    series[2].setData(data3);
}

function updateDistWithHRChart(series, index, remark) {
    var points0 = series[0].points;
    var points1 = series[1].points;
    var points2 = series[2].points;
    var ys0 = [];
    var ys1 = [];
    var ys2 = [];
    points0.forEach(function (p, i, arr) {
        ys0.push(p.y);
    });
    points1.forEach(function (p, i, arr) {
        ys1.push(p.y);
    });
    points2.forEach(function (p, i, arr) {
        ys2.push(p.y);
    });
    if (remark != null && remark.distwithhr != null && remark.distwithhr.length == 3 && ys0.length > index && ys1.length > index && ys2.length > index) {
        if (ys0[index] < Number(remark.distwithhr[0].toFixed(2))) {
            ys0[index] = Number(remark.distwithhr[0].toFixed(2));
            series[0].setData(ys0);
        }

        if (ys1[index] < Number(remark.distwithhr[1].toFixed(2))) {
            ys1[index] = Number(remark.distwithhr[1].toFixed(2));
            series[1].setData(ys1);
        }

        if (ys2[index] < Number(remark.distwithhr[2].toFixed(2))) {
            ys2[index] = Number(remark.distwithhr[2].toFixed(2));
            series[2].setData(ys2);
        }
    }
}
