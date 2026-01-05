<script lang="ts">
  import { formatDate, formatRelative, formatTime } from '$lib/src/temporal';

  type Props = {
    ms: number;
  };

  let { ms } = $props();
</script>

<div class="relative">
  <span class="peer">{formatRelative(ms)}</span>
  <div
    class="invisible absolute -top-full left-1/2 -translate-x-1/2 rounded-box bg-base-900 px-2 py-1 text-nowrap peer-hover:visible hover:visible">
    <span>
      {formatDate(ms)}
      <span class="text-primary">{formatTime(ms)}</span>
    </span>
  </div>
</div>
