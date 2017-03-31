import IndexBooks from './books/Index'
import ShowBook from './books/Show'

export default {
  routes: [
    {
      path: '/reading/books/pages/:page',
      name: 'reading.books.pages',
      component: IndexBooks
    },
    {
      path: '/reading/books/show/:id',
      name: 'reading.books.show',
      component: ShowBook
    }
  ],
  links: [
    {
      href: {name: 'reading.books.pages', params: {page: 1}},
      label: 'reading.books.index.title'
    }
  ],
  dashboard (user) {
    return null
  }
}
