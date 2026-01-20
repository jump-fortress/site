<script lang="ts">
  import { formatRunTime } from '$lib/src/times';
  import { Chart } from 'chart.js/auto';
  import 'chartjs-adapter-date-fns';

  import type { TimeWithPlayer } from '$lib/schema';

  type Props = {
    data: TimeWithPlayer[];
    div: string;
  };

  let { data, div }: Props = $props();

  const timeData = $derived(
    data.map(({ player, time }) => ({
      x: time.created_at,
      y: time.run_time
    }))
  );

  const divisionColors = new Map([
    ['Diamond', '#94e2d5'],
    ['Platinum', '#74c7ec'],
    ['Gold', '#f9e2af'],
    ['Silver', '#f5e0dc'],
    ['Bronze', '#fab387'],
    ['Steel', '#cdd6f4'],
    ['Wood', '#f2cdcd']
  ]);

  function chart(nd: Array<object>) {
    return (node: HTMLCanvasElement) => {
      const chart = new Chart(node, {
        type: 'scatter',
        data: {
          datasets: [{ data: nd }]
        },
        options: {
          font: {
            family: 'fredoka'
          },
          plugins: {
            legend: {
              display: false
            }
          },
          layout: {
            padding: 0
          },
          interaction: {
            mode: 'nearest'
          },
          aspectRatio: 4,
          elements: {
            line: {
              borderColor: divisionColors.get(div),
              borderWidth: 2
            },
            point: {
              pointStyle: 'circle',
              hitRadius: 1024,
              backgroundColor: divisionColors.get(div),
              hoverBorderColor: '#cdd6f4',
              hoverBorderWidth: 2,
              radius: 4,
              hoverRadius: 8
            }
          },
          scales: {
            x: {
              type: 'time',
              time: {
                unit: 'hour'
              },

              ticks: {
                stepSize: 1
              }
            },
            y: {
              ticks: {
                callback: function (value, _index, _ticks) {
                  return formatRunTime(parseInt(value.toString())).substring(0, 5);
                }
              }
            }
          }
        }
      });

      return () => {
        chart.destroy();
      };
    };
  }
</script>

<canvas {@attach chart(timeData)}></canvas>
