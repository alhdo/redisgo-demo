<template>
  <div class="card">
    <div class="card-header">
      <slot name="heading"></slot>
    </div>
    <div class="card-body">
      <div class="vlille-container">
        <div class="vlille-container-info">
          <div class="vlille-container-info-address"><slot name="address"></slot></div>
          <div class="vlille-container-info-city"><slot name="city"></slot></div>
          <div class="vlille-opening-hours">
            <span class="open">Ouvert</span>
          </div>
        </div>

        <div class="vlille-container-half-infos">
          <slot name="empty-space"></slot>
          <slot name="free-bikes"></slot>
        </div>
      </div>
    </div>
  </div>
</template>
