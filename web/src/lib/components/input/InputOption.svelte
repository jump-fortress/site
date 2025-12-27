<script lang="ts">
  // index & selectedIndex are used for the InputAutofill component
  type Props = {
    value: string;
    onSelect: (value: string) => void;
    index?: number | null;
    selectedIndex?: number | null;
  };

  let { value, onSelect, index = null, selectedIndex = null }: Props = $props();

  let option: HTMLButtonElement | undefined = $state();

  $effect(() => {
    if (index === selectedIndex) {
      option?.scrollIntoView({ behavior: 'smooth' });
    }
  });
</script>

<button
  bind:this={option}
  onmousedown={() => {
    onSelect(value);
  }}
  class="group flex w-full text-left text-base"
  ><span
    class="w-full px-2 group-hover:bg-jfgray-800
    {index === selectedIndex ? 'bg-jfgray-800' : ''}">{value}</span
  ></button>
