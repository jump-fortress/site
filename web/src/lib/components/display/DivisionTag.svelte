<script lang="ts">
  interface Props {
    div: string;
  }

  let { div }: Props = $props();
  let lowercaseDiv = $derived(div.toLowerCase());
</script>

<span
  class="rounded-full border-division-{lowercaseDiv} border-x-2 text-division-{lowercaseDiv} w-fit bg-base-900/50 px-2 pb-0.5 text-sm">
  {div}
</span>
