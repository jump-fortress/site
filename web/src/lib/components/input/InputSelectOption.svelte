<script lang="ts">
  type Props = {
    value: string;
    onSelect: (value: string) => void;
  };

  let { value, onSelect }: Props = $props();
</script>

<button
  onmousedown={() => {
    onSelect(value);
  }}
  class="group flex w-full text-left"
  ><span class="w-full rounded-lg px-2 group-hover:bg-jfgray-800">{value}</span></button
>
