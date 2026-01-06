<script lang="ts">
  import ClassImage from '../ClassImage.svelte';

  import type { Competition } from '$lib/schema';

  type Props = {
    competition: Competition;
    format: string;
    formatId: number;
  };

  let { competition, format, formatId }: Props = $props();
</script>

<div class="group flex h-8 cursor-pointer items-center gap-2 rounded-box px-2 hover:bg-base-900">
  <ClassImage selected={competition.class} />
  <span class="truncate group-hover:underline">{format} #{formatId}</span>
</div>
