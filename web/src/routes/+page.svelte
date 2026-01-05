<script lang="ts">
  import { formatDate, formatTime } from '$lib/src/temporal';
  import { Intl, Temporal } from 'temporal-polyfill';

  // input: time, in current timezone
  // output: time, in current timezone displaying utc offset

  const now = Temporal.Now.instant().epochMilliseconds;
</script>

<span>
  {formatDate(now)}
  <span class="text-primary">{formatTime(now)}</span>
</span>
