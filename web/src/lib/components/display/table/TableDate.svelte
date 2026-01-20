<script lang="ts">
  import { formatDate, formatRelative, formatTime } from '$lib/src/temporal';
  import { Temporal } from 'temporal-polyfill';

  type Props = {
    /** Format: date-time */
    date: string;
    fade?: boolean;
  };

  let { date, fade = true }: Props = $props();
  let ms = $derived(Temporal.Instant.from(date).epochMilliseconds);
</script>

<div class="relative">
  <span class="peer {fade ? 'text-content/75' : ''}">{formatRelative(ms)}</span>
  <div
    class="invisible absolute -top-1 right-0 rounded-box bg-base-900 px-2 py-1 text-nowrap drop-shadow-sm/100 drop-shadow-base-900 peer-hover:visible hover:visible">
    <span>
      {formatDate(ms)}
      <span class="text-primary">{formatTime(ms)}</span>
    </span>
  </div>
</div>
