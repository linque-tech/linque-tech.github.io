(function () {
  var vh = window.innerHeight;

  // Headings: dark navy → bright sky
  var hStart = { r: 45,  g: 90,  b: 138 };  // #2d5a8a
  var hEnd   = { r: 122, g: 213, b: 255 };  // #7ad5ff

  // Light-background sections: muted blue-grey → current light
  var sStart = { r: 90,  g: 105, b: 140 };  // #5a698c
  var sEnd   = { r: 217, g: 221, b: 235 };  // #d9ddeb

  var headings = Array.prototype.slice.call(document.querySelectorAll(
    '.heading-section-subtitle:not(.heading-hero-section):not(.jobs)'
  ));
  var sections = Array.prototype.slice.call(document.querySelectorAll(
    '.section.light_background'
  ));

  if (!headings.length && !sections.length) return;

  function lerp(a, b, t) { return Math.round(a + (b - a) * t); }
  function toRgb(s, e, t) {
    return 'rgb(' + lerp(s.r, e.r, t) + ',' + lerp(s.g, e.g, t) + ',' + lerp(s.b, e.b, t) + ')';
  }
  function scrollTarget(el) {
    var top = el.getBoundingClientRect().top;
    return Math.max(0, Math.min(1, (vh - top) / (vh / 2)));
  }

  // Each item starts dark (cur: 0) so visible elements animate in on page load
  var hItems = headings.map(function (h) { return { el: h, cur: 0, target: 0 }; });
  var sItems = sections.map(function (s) { return { el: s, cur: 0, target: 0 }; });

  // Smoothing factor — how quickly the displayed value chases the scroll target.
  // At 60 fps this gives roughly a 0.4-0.5 s ease-out feel.
  var SMOOTH = 0.07;

  function onScroll() {
    vh = window.innerHeight;
    hItems.forEach(function (h) { h.target = scrollTarget(h.el); });
    sItems.forEach(function (s) { s.target = scrollTarget(s.el); });
  }

  function tick() {
    hItems.forEach(function (h) {
      if (Math.abs(h.target - h.cur) > 0.0005) {
        h.cur += (h.target - h.cur) * SMOOTH;
        h.el.style.color = toRgb(hStart, hEnd, h.cur);
      }
    });
    sItems.forEach(function (s) {
      if (Math.abs(s.target - s.cur) > 0.0005) {
        s.cur += (s.target - s.cur) * SMOOTH;
        s.el.style.backgroundColor = toRgb(sStart, sEnd, s.cur);
      }
    });
    requestAnimationFrame(tick);
  }

  window.addEventListener('scroll', onScroll, { passive: true });
  window.addEventListener('resize', onScroll, { passive: true });
  onScroll();
  requestAnimationFrame(tick);
})();
