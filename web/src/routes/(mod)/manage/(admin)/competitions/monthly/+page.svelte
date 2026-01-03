<script lang="ts">
  import DivisionTag from '$lib/components/display/DivisionTag.svelte';
  import Input from '$lib/components/input/Input.svelte';
  import Label from '$lib/components/input/Label.svelte';
  import Select from '$lib/components/input/select/Select.svelte';
  import Section from '$lib/components/layout/Section.svelte';
  import { divisions } from '$lib/src/divisions.js';
  import { SvelteSet } from 'svelte/reactivity';

  import type { MonthlyInputBody } from '$lib/schema.js';

  let { data } = $props();

  // todo: prizes
  let competitionDivisions: SvelteSet<string> = new SvelteSet();
  let competitionMaps: SvelteSet<string> = new SvelteSet();
  let competitionClass = $state('Soldier');
  let competitionDate = $state();
  let competitionTime = $state();

  let a: MonthlyInputBody = $derived({
    competition: {
      class: 'Soldier',
      starts_at: 'a'
    },
    divisions: [
      {
        division: 'Platinum',
        map: 'a'
      }
    ]
  });
</script>

{#each competitionDivisions as div}
  <DivisionTag {div} />
{/each}

<Section label="create">
  <div class="flex gap-2">
    <Label label="start date" width={'max-w-60'}>
      <Input
        type="date"
        onsubmit={async (value) => {
          return null;
        }} />
    </Label>
    <Label label="start time" width={'max-w-60'}>
      <Input
        type="time"
        onsubmit={async (value) => {
          return null;
        }} />
    </Label>
  </div>

  {#await data.maps then maps}
    {#if maps}
      <Label label="add map">
        <Select
          type="text"
          options={maps}
          onsubmit={async (value) => {
            return null;
          }} />
      </Label>
    {/if}
  {/await}

  <Label label="add division">
    <Select
      type="text"
      options={divisions}
      onsubmit={async (value) => {
        competitionDivisions.add(value);
        return null;
      }} />
  </Label>
</Section>
