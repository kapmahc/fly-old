<template>
  <p v-html="marked(body)" />
</template>

<script>
import marked, { Renderer } from 'marked'
import highlightjs from 'highlight.js'

// Create your custom renderer.
const renderer = new Renderer()
renderer.code = (code, language) => {
  // Check whether the given language is valid for highlight.js.
  const validLang = !!(language && highlightjs.getLanguage(language))
  // Highlight only if the language is valid.
  const highlighted = validLang ? highlightjs.highlight(language, code).value : code
  // Render the highlighted code with `hljs` class.
  return `<pre><code class="hljs ${language}">${highlighted}</code></pre>`
}

// Set the renderer to marked.
marked.setOptions({ renderer })

export default {
  name: 'markdown-panel',
  props: ['body'],
  methods: {
    marked
  }
}
</script>
