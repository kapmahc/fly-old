$(document).ready(function() {
  $('div.markdown').each(function(i, block) {
    $(this).html(marked($(this).text()));    
  });
  $('pre code').each(function(i, block) {
    hljs.highlightBlock(block);
  });
});
