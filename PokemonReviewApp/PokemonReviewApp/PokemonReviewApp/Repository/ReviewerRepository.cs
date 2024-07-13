using AutoMapper;
using Microsoft.EntityFrameworkCore;
using PokemonReviewApp.Data;
using PokemonReviewApp.Interfaces;
using PokemonReviewApp.Models;

namespace PokemonReviewApp.Repository
{
    public class ReviewerRepository : IReviewerRepository
    {
        private readonly DataContext _context;
        private readonly IMapper _mapper;

        public ReviewerRepository(DataContext context, IMapper mapper)
        {
            _context = context;
            _mapper = mapper;
        }
        Reviewer IReviewerRepository.GetReviewer(int id)
        {
            return _context.Reviewers.Where(r => r.Id == id).Include(e => e.Reviews).FirstOrDefault();
        }

        ICollection<Reviewer> IReviewerRepository.GetReviewers()
        {
            return _context.Reviewers.ToList();
        }

        ICollection<Review> IReviewerRepository.GetReviewsByReviewer(int reviewerId)
        {
            return _context.Reviews.Where(r => r.Reviewer.Id == reviewerId).ToList();
        }

        bool IReviewerRepository.ReviewerExists(int id)
        {
            return _context.Reviewers.Any(r => r.Id == id);
        }
    }
}
