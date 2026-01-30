<script lang="ts">
  import { datetimeToMs, formatDate, formatRelative, formatTime } from '$lib/helpers/temporal';

  type Props = {
    datetime: string;
  };

  let { datetime }: Props = $props();
  const ms = $derived(datetimeToMs(datetime));
</script>

<div class="relative">
  <span class="peer">{formatRelative(ms)}</span>
  <div
    class="invisible absolute -top-1 right-0 rounded-box bg-base-900 px-2 py-1 text-nowrap text-content drop-shadow-sm/100 drop-shadow-base-900 peer-hover:visible hover:visible">
    <span>
      {formatDate(ms)}
      <span class="text-primary">{formatTime(ms)}</span>
    </span>
  </div>
</div>
