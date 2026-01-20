export function formatRunTime(run_time: number) {
  const minutes = Math.floor(run_time / 60)
    .toString()
    .padStart(2, '0');
  const seconds = (run_time % 60).toFixed(3).padStart(6, '0');

  return `${minutes}:${seconds}`;
}
