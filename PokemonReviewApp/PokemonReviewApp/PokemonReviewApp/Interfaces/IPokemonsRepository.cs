using PokemonReviewApp.Models;

namespace PokemonReviewApp.Interfaces
{
    public interface IPokemonsRepository
    {
        ICollection<Pokemon> GetPokemons();
        Pokemon GetPokemon(int id);
        Pokemon GetPokemon(string name);
        decimal GetPokemonRatings(int pokeId);
        bool PokemonExists(int pokeId);
        bool CreatePokemon(int ownerId, int categoryId, Pokemon pokemon);
        bool Save();
    }
}
