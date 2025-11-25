<script lang="ts">
  type Props = {
    title: string;
    options: {
      src: string;
      value: string;
    }[];
    selectedOption: string;
    onSelect: (value: string) => void;
  };

  let { title, options, selectedOption = $bindable(), onSelect }: Props = $props();
</script>

<div class="h-18 border-jfgray-700 relative mt-2 flex w-fit text-nowrap border-2">
  <span class="bg-jfgray-800 leading-1 absolute -top-1 left-2 z-10 px-1 text-base">{title}</span>
  {#each options as option}
    <button
      onclick={() => {
        onSelect(option.value);
      }}
      class="{selectedOption === option.value
        ? 'bg-jfgray-900 opacity-100'
        : 'opacity-50'} flex w-20 cursor-pointer justify-center transition-opacity hover:opacity-100"
    >
      <img
        src={option.src}
        alt=""
        draggable="false"
        class="{option.value === 'None' ? 'object-none' : 'object-contain'} pb-2 pt-3"
      />
    </button>
  {/each}
</div>

<style>
  img {
    filter: brightness(0) saturate(100%) invert(97%) sepia(49%) saturate(6023%) hue-rotate(179deg)
      brightness(100%) contrast(108%);
  }
</style>
