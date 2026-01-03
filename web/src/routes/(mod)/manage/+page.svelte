<script lang="ts">
  import Section from '$lib/components/layout/Section.svelte';
</script>

<Section label="moderation">
  <div class="flex flex-col">
    <h2 class="text-lg text-primary">moderation pages are split among 4 roles:</h2>
    <span>admins - can manage competitions, players, and payouts.</span>
    <span>moderators - can manage players.</span>
    <span>consultants - can view detailed info for players.</span>
    <span>
      treasurers - can view detailed info for players, view pending payouts for players, and send
      automated trade links.
    </span>
  </div>
</Section>
