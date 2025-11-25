<script lang="ts">
  interface Props {
    div: string;
    styles?: string;
  }

  let { div, styles = '' }: Props = $props();
</script>

<span
  class="rounded-full px-2 pb-0.5 border-division-{div} border-x-3 text-division-{div} bg-jfgray-900/75 {styles} text-base"
  >{div}</span
>
