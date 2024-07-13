﻿using PokemonReviewApp.Models;

namespace PokemonReviewApp.Interfaces
{
    public interface ICategoryRepository
    {
        ICollection<Category> GetCategories();
        Category GetCategory(int id);

        ICollection<Pokemon> GetPokemonsByCategory(int CategoryId);
        bool CategoriesExists(int id);

    }
}
