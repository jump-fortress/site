<script lang="ts">
  import { divisions } from '$lib/divisions';

  interface Props {
    div: string;
    styles?: string;
  }

  let { div, styles = '' }: Props = $props();
  let caseInsensitiveDiv = $derived(div.toLowerCase());
</script>

<span
  class="rounded-full px-2 pb-0.5 border-division-{caseInsensitiveDiv} border-x-3 text-division-{caseInsensitiveDiv} bg-jfgray-900/75 {styles} w-fit text-base"
  >{div}</span>
