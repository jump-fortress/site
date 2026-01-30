<script lang="ts">
  import ClassImage from '../ClassImage.svelte';

  import type { Event } from '$lib/schema';

  type Props = {
    event: Event;
  };

  let { event }: Props = $props();
</script>

<div class="group flex h-8 cursor-pointer items-center gap-2 rounded-box px-2 hover:bg-base-900">
  <ClassImage player_class={event.player_class} />
  <span class="truncate group-hover:underline">{event.kind} #{event.kind_id}</span>
</div>
