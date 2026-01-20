<script lang="ts">
  import { dateToMs } from '$lib/src/temporal';
  import { formatRunTime } from '$lib/src/times';
  import { Chart } from 'chart.js/auto';
  import 'chartjs-adapter-date-fns';
  import { onMount } from 'svelte';
  import { Temporal } from 'temporal-polyfill';

  import type { TimeWithPlayer } from '$lib/schema';
  import type { Attachment } from 'svelte/attachments';

  type Props = {
    data: TimeWithPlayer[];
  };

  let { data }: Props = $props();

  $inspect(dateToMs(data.at(0)?.time.created_at ?? ''));

  const timeData = $derived(data.map(({ time }) => ({ x: time.created_at, y: time.run_time })));

  function chart(nd: Array<object>) {
    return (node: HTMLCanvasElement) => {
      const chart = new Chart(node, {
        type: 'scatter',
        data: {
          datasets: [{ data: nd }]
        },

        options: {
          scales: {
            x: {
              type: 'time',
              time: {
                unit: 'hour'
              },
              title: {
                display: true,
                text: 'run time'
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
