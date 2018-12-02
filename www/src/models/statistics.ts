import ξ from 'mithril';

export const default {
    const chartColors = {
        red: 'rgb(255, 99, 132)',
        orange: 'rgb(255, 159, 64)',
        yellow: 'rgb(255, 205, 86)',
        green: 'rgb(75, 192, 192)',
        blue: 'rgb(54, 162, 235)',
        purple: 'rgb(153, 102, 255)',
        grey: 'rgb(201, 203, 207)',
    };

    const MONTHS = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'];
    const config = {
        type: 'line',
        data: {
            labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July'],
            datasets: [{
                label: 'My first dataset!',
                backgroundColor: this.chartColors.red,
                borderColor: this.chartColors.red,
                data: [
                    100,
                    10,
                    70,
                    -10,
                    -90,
                    30,
                    5,
                ],
                fill: false,
            }, {
                label: 'My first dataset!',
                backgroundColor: this.chartColors.blue,
                borderColor: this.chartColors.blue,
                data: [
                    100,
                    50,
                    50,
                    10,
                    -70,
                    -100,
                    45,
                ],
                fill: true,
            }],
        },
    };


    let options: {
        responsive: true,
        title: {
            display: true,
            text: 'Chart.js Line Chart'
        },
        tooltips: {
            mode: 'index',
            intersect: false,
        },
        hover: {
            mode: 'nearest',
            intersect: true
        },
        scales: {
            xAxes: [{
                display: true,
                scaleLabel: {
                    display: true,
                    labelString: 'Month'
                }
            }],
            yAxes: [{
                display: true,
                scaleLabel: {
                    display: true,
                    labelString: 'Value'
                }
            }]
        }
    };
};
