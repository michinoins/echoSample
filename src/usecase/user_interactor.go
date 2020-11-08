package usecase

import "echoSample/src/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) {
	interactor.UserRepository.Store(u)
}

func (interactor *UserInteractor) GetInfo() []domain.User {
	return interactor.UserRepository.Select()
}

func (interactor *UserInteractor) Delete(id string) {
	interactor.UserRepository.Delete(id)
}

//func (interactor *UserInteractor) Add(u domain.User)という書き方で、関数内でinteractorを使うと、それはUserRepositoryを使っているということになるっポイ?(その中で関数を読んだりしてたら継承?s)
