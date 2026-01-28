<script lang="ts">
  import Label from '$lib/components/input/Label.svelte';
  import { fade } from 'svelte/transition';
  import SubmitIcon from '$lib/components/input/SubmitIcon.svelte';

  type Props = {
    label?: string;
    type: string;
    max_width?: string;
    placeholder?: string;
    value?: string;
    onsubmit: (value: string) => Promise<boolean>;
  };

  let {
    label = '',
    type,
    max_width = 'max-w-80',
    placeholder = '',
    value = $bindable(),
    onsubmit
  }: Props = $props();

  let valid: Promise<boolean> = $state(Promise.resolve(true));
</script>

<Label {label} {max_width}>
  <input
    class="size-full px-2 text-primary"
    {type}
    {placeholder}
    bind:value
    onkeydown={async (e) => {
      if (e.key === 'Enter' && value) {
        valid = onsubmit(value);
        if (await valid) {
          placeholder = value;
          value = '';
        }
      }
    }} />
  <!-- submit button & response -->
  <button
    in:fade
    class="relative grid size-9 cursor-pointer place-content-center text-content"
    tabindex="-1"
    onclick={async () => {
      if (value) {
        valid = onsubmit(value);
        if (await valid) {
          placeholder = value;
          value = '';
        }
      }
    }}>
    <SubmitIcon {valid} />
  </button>
</Label>
