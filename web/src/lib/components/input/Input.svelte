<script lang="ts">
  import Response from '$lib/components/input/Response.svelte';
  import SubmitIcon from '$lib/components/input/SubmitIcon.svelte';
  import { fade } from 'svelte/transition';

  type Props = {
    type: string;
    placeholder?: string;
    message?: string;
    clear?: boolean;
    onsubmit: (value: string) => Promise<InputResponse>;
  };

  let { type, placeholder = '', message = '', clear = false, onsubmit }: Props = $props();

  let value: string = $state('');
  let response: Promise<InputResponse> = $derived(
    Promise.resolve({ error: false, message: message })
  );
</script>

<input
  class="size-full px-2 text-primary"
  {type}
  {placeholder}
  bind:value
  onkeydown={(event) => {
    if (event.key === 'Enter' && value) {
      response = onsubmit(value);
      if (clear) {
        value = '';
      }
    }
  }} />

<!-- submit button & response -->
<button
  in:fade
  class="relative grid size-10 cursor-pointer place-content-center text-content"
  tabindex="-1"
  onclick={() => {
    if (value) {
      response = onsubmit(value);
      if (clear) {
        value = '';
      }
    }
  }}>
  <SubmitIcon {response} />
</button>
<Response {response} />
