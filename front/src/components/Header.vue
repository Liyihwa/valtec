<template>
  <div class="py-[0.75rem] px-[1.5rem] border-b-1 border-b-[#323232] w-full">
    <div class="flex items-center justify-between max-w-[64rem] mx-auto h-full">
      <div class="text-2xl fw-800 text-gradient text-transparent">
        <a href="/">ValorantTechLab</a>
      </div>
      <div class="flex justify-between items-center gap-7">
        <div class="hidden sm:inline">
          <div class="flex gap-4">
            <Modules />
          </div>
        </div>
        <ThemeToggle />
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup></script>
