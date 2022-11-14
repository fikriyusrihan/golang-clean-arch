package registry

import (
	"github.com/fikriyusrihan/golang-clean-arch/interface/controller"
	ip "github.com/fikriyusrihan/golang-clean-arch/interface/presenter"
	ir "github.com/fikriyusrihan/golang-clean-arch/interface/repository"
	"github.com/fikriyusrihan/golang-clean-arch/usecase/interactor"
	up "github.com/fikriyusrihan/golang-clean-arch/usecase/presenter"
	ur "github.com/fikriyusrihan/golang-clean-arch/usecase/repository"
)

func (r *registry) NewBookController() controller.BookController {
	return controller.NewBookController(r.NewBookInteractor())
}

func (r *registry) NewBookInteractor() interactor.BookInteractor {
	return interactor.NewBookInteractor(r.NewBookRepository(), r.NewBookPresenter())
}

func (r *registry) NewBookRepository() ur.BookRepository {
	return ir.NewBookRepository(r.db)
}

func (r *registry) NewBookPresenter() up.BookPresenter {
	return ip.NewBookPresenter()
}
